package main

import (
	"fmt"
	"log"
	"time"
	"flag"
	"strings"
	"math/rand"

	"github.com/nats-io/nuid"
	"github.com/garyburd/redigo/redis"
)

type strsli []string

func (s *strsli) String() string {
	return fmt.Sprint(*s)
}

func (s *strsli) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

var tv, date string
var adUuids strsli

func init() {
	// Two flags sharing a variable, so we can have a shorthand.
	// The order of initialization is undefined, so make sure both use the
	// same default value. They must be set up with an init function.
	flag.StringVar(&tv, "tv", "liaoningTV", "The tv name.")
	flag.StringVar(&tv, "t", "liaoningTV", "The tv name.")
	flag.StringVar(&date, "date", "2017-05-27", "The date.")
	flag.StringVar(&date, "d", "2017-05-27", "The date.")

	// Tie the command-line flag to the adUuids variable and
	// set a usage message.
	flag.Var(&adUuids, "uids", "comma-separatd list of strsli to use")
}

func main() {
	flag.Parse()

	// adUuids = []string{"a56209a9-082b-e32e-043d-213a3e669df5"}
	// tv = "liaoningTV"
	// date = "2017-05-27"

	theTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatalf("The parameter date: %s is not valid", date)
	}
	timestamp := theTime.Unix()

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("Connect redis failed: %v", err)
	}
	defer c.Close()

	rkey := "cc:raw:"+tv+":"+theTime.Format("0102")
	fmt.Println(rkey)

	curtime := timestamp
	for curtime <= timestamp + 86400 {
		rand.Seed(time.Now().UnixNano())
		curtime = curtime + int64(rand.Intn(10)) + 1

		for _, adUuid := range adUuids {
			uuid := nuid.Next()
			member := adUuid+"-"+uuid
			fmt.Println(member, curtime)
			if _, err := c.Do("ZADD", rkey, curtime, member); err != nil {
				log.Fatalf("%v", err)
			}
		}
	}
}
