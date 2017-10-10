package main

import (
	"fmt"
	"sync"
)

var global int
var wg sync.WaitGroup
var w sync.Mutex

func count() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		w.Lock()
		global++
		w.Unlock()
	}
}

func main() {
	wg.Add(2)
	go count()
	go count()
	wg.Wait()
	fmt.Println(global)
}
