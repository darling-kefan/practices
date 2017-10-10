package main

import (
	"fmt"
	_ "math"
)

func main() {
	shows, m, n := 6170, 0, 0
	fmt.Println(float64(shows) / (float64(m)*1.2 + float64(n)))
}
