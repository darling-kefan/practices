package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = []string{"I", "make", "a", "mistake", "on", "last", "sunday"}
	b := []string(a)[1:]
	fmt.Println(reflect.TypeOf(b), b)
}
