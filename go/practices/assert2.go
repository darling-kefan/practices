package main

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct{}

func (car) String() string {
	return "Vroom!"
}

type cloud struct{}

func (cloud) String() string {
	return "Big Data!"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	for i := 0; i < 10; i++ {
		index := rand.Intn(2)

		if v, is := mvs[index].(cloud); is {
			fmt.Println("Go Lucky:", v)
			continue
		}

		fmt.Println("Got Unlucky")
	}
}
