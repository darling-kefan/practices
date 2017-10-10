package main

import (
	"fmt"
)

func yanghui(n int64) []int64 {
	var l []int64
	var i int64
	for i = 1; i <= n; i++ {
		if i == 1 {
			l = []int64{1}
		} else if i == 2 {
			l = []int64{1, 1}
		} else {
			tmp := make([]int64, i)
			tmp[0] = 1
			tmp[i-1] = 1
			for j := 0; j < len(l)-1; j++ {
				tmp[j+1] = l[j] + l[j+1]
			}
			l = tmp
		}
		fmt.Println(l)
	}
	return l
}

func main() {
	l := yanghui(20)
	fmt.Println(l)
}
