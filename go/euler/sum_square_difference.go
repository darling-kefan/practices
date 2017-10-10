package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing argument")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sqs := squareSum(n)
	sus := sumSquare(n)

	fmt.Println(sqs, sus, sus-sqs)
}

func squareSum(n int) (sum int) {
	if n <= 0 {
		return
	}

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return
}

func sumSquare(n int) (squ int) {
	if n <= 0 {
		return
	}

	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}
