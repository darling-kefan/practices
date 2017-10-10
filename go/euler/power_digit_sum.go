package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var i int64
	mul := big.NewInt(1)
	for i = 1; i <= 1000; i++ {
		mul = mul.Mul(mul, big.NewInt(2))
	}
	fmt.Println(mul)

	sum := 0
	for _, v := range mul.String() {
		if vv, err := strconv.Atoi(string(v)); err == nil {
			sum += vv
		}
	}
	fmt.Println(sum)
}
