package main

import "fmt"

func main() {
	// Create a slice with a length of 5 elements and a capacity of 8.
	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	inspectSlice(slice)
}

func inspectSlice(slice []string) {
	fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("The index %d of slice address %p and value %s\n",
			i, &slice[i], slice[i])
	}
}
