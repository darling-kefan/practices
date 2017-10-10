package main

import "fmt"

func main() {
	// Declare a slice of integers with 7 values.
	x := make([]int, 7)

	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}

	twohundred := &x[1]

	x = append(x, 800)

	x[1]++

	fmt.Println("Address1:", twohundred, "Pointer:", *twohundred, "Address2:", &x[1], "Element", x[1])
}
