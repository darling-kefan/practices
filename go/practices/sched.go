package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	a := 1
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		a = a + 1
	}()

	go func() {
		defer wg.Done()
		if a == 1 {
			runtime.Gosched()
			fmt.Println("a==", a)
		}
	}()

	runtime.Gosched()
	wg.Wait()
}
