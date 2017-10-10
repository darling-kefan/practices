package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing parameter.")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i, pc := 2, 0
	for {
		if isPrime(i) {
			pc++
		}
		if pc >= n {
			break
		}
		i += 1
	}

	fmt.Printf("The %dst prime number is %d\n", n, i)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
