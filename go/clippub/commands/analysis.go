package main

import (
	"fmt"
	"log"
	"time"
	"sort"
	"flag"
	"strings"
	"strconv"

	"github.com/nats-io/nuid"
	"github.com/garyburd/redigo/redis"
)

// 最小时间间隔：用于区分同一广告是否为同一次播放
const AD_MINIMUM_INTERVAL time.Duration = 10

var rawKey, dstKey string

var tv string

func init() {
	flag.StringVar(&tv, "tv", "liaoningTV", "The tv name.")
	flag.StringVar(&tv, "t", "liaoningTV", "The tv name.")
}

func main() {
	flag.Parse()

	redisConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("Connect redis failed: %v", err)
	}
	defer redisConn.Close()

	rawKey = "cc:raw:"+tv+":"+time.Now().Format("0102")
	dstKey = "cc:"+tv+":"+time.Now().Format("0102")
	fmt.Println(rawKey, dstKey)
	
	for {
		processor(redisConn)
		time.Sleep(20 * time.Second)
	}
}

func processor(redisConn redis.Conn) {
	timestamp := time.Now().Unix()

	adUuid2Points := make(map[string][]int)
	rawSet, err := redis.Strings(redisConn.Do("ZRANGEBYSCORE", rawKey, timestamp - 30, timestamp, "withscores"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	if len(rawSet) > 0 {
		for n := 0; n < len(rawSet); n = n + 2 {
			member := rawSet[n]
			score, err  := strconv.Atoi(rawSet[n+1])
			if err != nil {
				log.Fatalf("%v", score)
			}
			adUuid := member[:strings.LastIndex(member, "-")]
			if err != nil {
				log.Fatalf("%v", err)
			}
			adUuid2Points[adUuid] = append(adUuid2Points[adUuid], score)
		}
		// sort []int
		for _, v := range adUuid2Points {
			sort.Ints(v)
		}
	}

	if len(adUuid2Points) == 0 {
		return
	}
	
	dstSet, err := redis.Strings(redisConn.Do("ZREVRANGEBYSCORE", dstKey, timestamp, timestamp - 60, "withscores"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	/*dstMap := make(map[string]int)
	if (len(dstSet) > 0) {
		for n := 0; n < len(dstSet); n += 2 {
			member := dstSet[n]
			score, err := strconv.Atoi(dstSet[n+1])
			if err != nil {
				log.Fatalf("%v", score)
			}
			dstMap[member] = score
		}
	}*/

	fmt.Printf("%#v, %#v\n", adUuid2Points, dstSet)

	
	// 测试数据
	/*adUuid2Points = map[string][]int{
		"a56209a9-082b-e32e-043d-213a3e669df5": []int{
			1495708171,
			1495708177,
			1495708179,
			1495708188,
			1495708198,
			1495708200,
		},
	}*/

	
	for adUuid, points := range adUuid2Points {
		// 判断广告adUuid是否在播放中
		isRunning := false
		// 定义播放中广告{adUuid}-{mark}-e
		eMember := ""
		for n := 0; n < len(dstSet); n = n + 2 {
			adUuidMark := dstSet[n];
			score, err := strconv.Atoi(dstSet[n+1])
			if err != nil {
				log.Fatalf("%v", err)
			}

			if strings.Index(adUuidMark, adUuid) != -1 && len(strings.TrimRight(adUuidMark, "-e")) != len(adUuidMark) {
				// 清理掉集合里小于score的元素
				for idx, point := range points {
					if point > score {
						points = points[idx:]
						break
					}
				}

				// 如果rawscore-score小于10s，则认为广告adUuid在播放中；否则不在播放中
				if points[0] - score < int(AD_MINIMUM_INTERVAL) {
					isRunning = true
					eMember   = adUuidMark
				}

				break
			}
		}

		// @TODO: It's very important!!!
		// 由于Golang map每次遍历的顺序是随机的，每次遍历的顺序都不相同。因此，此处使用hash是错误的，必须使用Slice
		/*
		for adUuidMark, score := range dstMap {
			if strings.Index(adUuidMark, adUuid) != -1 && len(strings.TrimRight(adUuidMark, "-e")) != len(adUuidMark) {
				// 清理掉集合里小于score的元素
				for idx, point := range points {
					if point > score {
						points = points[idx:]
						break
					}
				}

				// 如果rawscore-score小于10s，则认为广告adUuid在播放中；否则不在播放中
				if points[0] - score < int(AD_MINIMUM_INTERVAL) {
					isRunning = true
					eMember   = adUuidMark
				}

				break
			}
		}*/

		fmt.Println(isRunning, points)
		
		if isRunning {
			ePoint := points[0]
			isUpdatedEnd := false
			for idx, point := range points {
				// 如果两个点位的间隔时间超过，则定为一个新广告开始；
				if point - ePoint < 10 {
					ePoint = point
				} else {
					// 更新广告单里播放中广告的结束时间
					if _, err := redisConn.Do("ZADD", dstKey, ePoint, eMember); err != nil {
						log.Fatalf("%v", err)
					}
					isUpdatedEnd = true
					appendAdList(adUuid, points[idx:], redisConn)
				}
			}
			if !isUpdatedEnd {
				// 更新广告单里播放中广告的结束时间
				if _, err := redisConn.Do("ZADD", dstKey, ePoint, eMember); err != nil {
					log.Fatalf("%v", err)
				}
				isUpdatedEnd = true
			}
		} else {
			appendAdList(adUuid, points, redisConn)
		}
	}
}

// 追加新广告至广告单
func appendAdList(adUuid string, points []int, redisConn redis.Conn) {
	if len(points) == 0 {
		return
	}
	
	mark := nuid.Next()
	sMember := adUuid+"-"+mark+"-s"
	eMember := adUuid+"-"+mark+"-e"
	sPoint := points[0]
	ePoint := points[0]
	
	// 将广告开始时间写入cc集合
	if _, err := redisConn.Do("ZADD", dstKey, sPoint, sMember); err != nil {
		log.Fatalf("%v", err)
	}

	isUpdatedEnd := false
	for idx, point := range points {
		// fmt.Println(idx, point, ePoint)
		if point - ePoint < 0 {
			continue
		} else if point - ePoint < 10 {
			ePoint = point
		} else {
			// fmt.Println(idx, point, ePoint)
			// 广告结束，更新前一个广告的结束时间
			if _, err := redisConn.Do("ZADD", dstKey, ePoint, eMember); err != nil {
				log.Fatalf("%v", err)
			}
			isUpdatedEnd = true
			appendAdList(adUuid, points[idx:], redisConn)
			// Notice: 此处必须加break,否则错误；错误测试数据如：
			/*adUuid2Points = map[string][]int{
				"a56209a9-082b-e32e-043d-213a3e669df5": []int{
					1495708171,
					1495708177,
					1495708179,
					1495708188,
					1495708198,
					1495708200,
				},
			}*/
			break
		}
	}

	if !isUpdatedEnd {
		// 广告结束，更新前一个广告的结束时间
		if _, err := redisConn.Do("ZADD", dstKey, ePoint, eMember); err != nil {
			log.Fatalf("%v", err)
		}
		isUpdatedEnd = true
	}
}
