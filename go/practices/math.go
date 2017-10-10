package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float32 = 3.0
	var y int = 10
	m := x * float32(y)
	n := x / float32(y)
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(m), m, reflect.TypeOf(n), n)
}
