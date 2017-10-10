package main

import (
	"fmt"
)

func addTail(s []int) {
	var ns [][]int
	for _, v := range []int{1, 2} {
		ns = append(ns, append(s, v))
	}
	fmt.Println(ns)
}

func main() {
	s1 := []int{0, 0}
	s2 := append(s1, 0)

	fmt.Printf("s1: length-%d, capacity-%d\n", len(s1), cap(s1))
	fmt.Printf("s2: length-%d, capacity-%d\n", len(s2), cap(s2))
	
	for _, v := range [][]int{s1, s2} {
		addTail(v)
	}
}

// output
// 
