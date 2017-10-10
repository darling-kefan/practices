package main

import "fmt"

func main() {
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// Parameters are [start_index : (start_index + length)]
	slice2 := slice1[2:4]
	inspectSlice(slice2)

	slice2[0] = "tangshouqiang"
	inspectSlice(slice1)
	inspectSlice(slice2)
}

func inspectSlice(slice []string) {
	fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))

	for i := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			slice[i])
	}
}
