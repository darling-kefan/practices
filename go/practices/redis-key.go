package main

import (
	"fmt"
	"log"
	"github.com/garyburd/redigo/redis"
)

func main() {
	redisConn, err := redis.Dial("tcp", "127.0.0.1:6311")
	if err != nil {
		log.Fatalf("Connect redis failed: %v", err)
	}
	defer redisConn.Close()

	var hours  []string = []string{"15", "16"}
	var agents []string = []string{"1", "2"}
	var ages   []string = []string{"0", "1", "2", "3", "4", "5", "6", "7"}

	member := "1-16-12-73622"
	score  := 100
	for _, hour := range hours {
		for _, agent := range agents {
			tcKey := "tc:"+agent+":"+hour+"367717"

			if _, err := redisConn.Do("ZADD", tcKey, score, member); err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Printf("zadd %s %s:%d\n", tcKey, member, score)
			
			for _, age := range ages {
				toKey := "to:"+agent+":"+hour+age+"00"
				if _, err := redisConn.Do("ZADD", toKey, score, member); err != nil {
					log.Fatalf("%v", err)
				}
				fmt.Printf("zadd %s %s:%d\n", toKey, member, score)
			}
		}
	}
}
