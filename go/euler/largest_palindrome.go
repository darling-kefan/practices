package main

import (
	"fmt"
	"strconv"
)

func main() {
	var maxMulti int
	for m := 999; m >= 100; m-- {
		for n := 999; n >= 100; n-- {
			multi := m * n
			if maxMulti > multi {
				break
			}
			if strconv.Itoa(multi) == reverseString(strconv.Itoa(multi)) && multi > maxMulti {
				maxMulti = multi
				fmt.Println(m, n, multi)
			}
		}
	}

	fmt.Println(maxMulti)
}

func reverseString(s string) string {
	bs := []rune(s)
	for from, to := 0, len(bs)-1; from < to; from, to = from+1, to-1 {
		bs[from], bs[to] = bs[to], bs[from]
	}
	return string(bs)
}
