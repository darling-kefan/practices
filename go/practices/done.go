package main

import (
	"fmt"
)

func main() {
	// Create a channel to synchronize goroutines
	done := make(chan bool)

	go func() {
		fmt.Println("goroutine message")
		// Tell the main function everything is done.
		// This channel is visible inside this goroutine because
		// it is executed in the same address space.
		done <- true
	}()

	fmt.Println("main message")
	<-done // Wait for the goroutine to finish
}
