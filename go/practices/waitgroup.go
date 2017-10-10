package main

import (
	"fmt"
	"sync"
)

func someRoutine(s string, wg *sync.WaitGroup) {
	fmt.Printf("Hey! Did you know the value is %s?\n", s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var n int = 7
	wg.Add(n)

	for i := 0; i < 7; i++ {
		go someRoutine("someValue", &wg)
	}

	wg.Wait() // this waits for the counter to be 0
	fmt.Println("All finished!")
}
