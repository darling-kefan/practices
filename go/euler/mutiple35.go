package main

import (
	"flag"
	"fmt"
)

var num int

func init() {
	flag.IntVar(&num, "n", 0, "number")
	flag.Parse()
}

func main() {
	fmt.Println(mutiple35(num))
}

func mutiple35(n int) (sum int) {
	if num < 3 {
		return
	}

	for i := 3; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return
}
