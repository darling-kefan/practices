package main

import "fmt"

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	// b()
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
