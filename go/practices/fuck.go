package main

import "fmt"

func main() {
	testMap := make(map[int64]int64)
	fmt.Println(testMap)
	testMap1 := FuckMap(testMap)
	fmt.Println(testMap, testMap1)
	testMap2 := FuckMap2(testMap)
	fmt.Println(testMap, testMap2)

	testSlice := []int64{}
	fmt.Println(testSlice)
	FuckSlice(testSlice)
	fmt.Println(testSlice)
	FuckSlice2(&testSlice)
	fmt.Println(testSlice)
}

func FuckMap(t map[int64]int64) map[int64]int64 {
	t[1] = 1
	return t
}

func FuckMap2(t map[int64]int64) map[int64]int64 {
	t[1] = 2
	return t
}

func FuckSlice(a []int64) []int64 {
	a = append(a, 1)
	return a
}

func FuckSlice2(a *[]int64) *[]int64 {
	*a = append(*a, 1)
	return a
}
