package main

import "fmt"

func main() {
	// Using the value semantic form of the for range
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for _, v := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", v)
	}

	//
	five = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", five[i])
	}
}
