package main

import "fmt"

//func main() {
//	// Declare a nil slice of strings.
//	var data []string
//
//	// Capture the capacity of the slice.
//	lastCap := cap(data)
//
//	for record := 1; record <= 102400; record++ {
//		// Use the build-in function append to add to the slice.
//		data = append(data, fmt.Sprintf("Rec: %d", record))
//
//		if lastCap != cap(data) {
//
//			prec := float64(cap(data)-lastCap) / float64(cap(data)) * 100
//
//			lastCap = cap(data)
//
//			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
//				&data[0],
//				record,
//				lastCap,
//				prec)
//		}
//	}
//}

func main() {
	// Declare a nil slice of strings.
	var data []int

	// Capture the capacity of the slice.
	lastCap := cap(data)

	for record := 1; record <= 102400; record++ {
		// Use the build-in function append to add to the slice.
		data = append(data, record)

		if lastCap != cap(data) {

			prec := float64(cap(data)-lastCap) / float64(cap(data)) * 100

			lastCap = cap(data)

			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				lastCap,
				prec)
		}
	}
}
