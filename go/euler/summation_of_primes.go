package main

import (
	"fmt"
	"math"
)

func main() {
	//var primes []int
	//for i := 2; i < 2000000; i++ {
	//	if isPrime(i) {
	//		primes = append(primes, i)
	//	}
	//}
	//
	//var sum int64
	//for _, v := range primes {
	//	sum += int64(v)
	//}
	//
	//fmt.Println(len(primes), sum)

	var sum int64
	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += int64(i)
		}
	}
	fmt.Println(sum)
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
