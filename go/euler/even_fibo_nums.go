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
	fiboSet := fibo(num)
	fiboSet = fiboSet[1:]
	fmt.Println(fiboSet)

	var evenSum int
	for _, v := range fiboSet {
		if v%2 == 0 {
			evenSum += v
		}
	}

	fmt.Println(evenSum)
}

func fibo(n int) (ret []int) {
	if n <= 0 {
		return
	}
	a, b := 0, 1
	ret = []int{a}
	for b < n {
		ret = append(ret, b)
		a, b = b, a+b
	}
	return
}
