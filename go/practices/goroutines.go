package main

import (
	"fmt"
	"time"
)

func makeMessage(s string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("I wanted to tell you '%s' for the %dth time", s, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func makeMessage2(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("I wanted to tell you '%s' for the %dth time", s, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string) // create channel

	go makeMessage2("Hello", ch)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	// makeMessage("Hello World")
}
