package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", 'x', 'x')
	fmt.Printf("%T %v\n", 0i, 0i)

	var u uint
	var v = -1
	u = uint(v)
	fmt.Printf("%T %v %T %v\n", u, u, v, v)

	var f float32 = 1
	var i int = 1.000
	var w uint32 = 1e3 - 99.0*10.0 - 9
	var c float64 = '\x01'
	var p uintptr = '\u0001'
	var r complex64 = 'b' - 'a'
	var b byte = 1.0 + 3i - 3.0i

	fmt.Println(f, i, w, c, p, r, b)

	var ff = 'a' * 1.5
	fmt.Printf("%T, %v\n", ff, ff)
}
