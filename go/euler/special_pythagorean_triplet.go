package main

import (
	"fmt"
)

func main() {
	a, b, c := 0, 0, 0
	for a = 1; a < 1000/3; a++ {
		for b = a + 1; b < 1000/2; b++ {
			c = 1000 - a - b
			if c <= b {
				break
			}
			if c*c == a*a+b*b {
				fmt.Println(a, b, c)
			}
		}
	}
}
