package main

import (
	"fmt"
	// "time"
)

/*
func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(1 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	// Output: 3 3 3
	// Output: 0 1 2
}

func main() {
	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
*/

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(v string) {
			fmt.Println(v)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before existing
	for _ = range values {
		<-done
	}
}
