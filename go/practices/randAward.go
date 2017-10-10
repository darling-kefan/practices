package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// map中，key代表名称，value代表成交单数
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 6,
		"c": 3,
		"d": 12,
		"e": 20,
		"f": 1,
	}

	// 记录随机中奖
	type ua struct {
		name  string
		count int64
	}
	i := 0
	uws := make([]ua, len(users))
	for u := range users {
		uws[i] = ua{name: u}
		i++
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < 10000; i++ {
		rdu := rand.Intn(len(uws))
		uws[rdu].count++
	}

	// 打印随机数
	for _, v := range uws {
		fmt.Printf("%s:\t%d\n", v.name, v.count)
	}
}
