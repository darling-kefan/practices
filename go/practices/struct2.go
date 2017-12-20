package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e1 := example{true, 12, 11.2}
	e2 := example{
		flag:    true,
		counter: 234,
		pi:      889.4,
	}
	var e3 example
	e4 := new(example)
	e5 := struct {
		flag    bool
		counter int16
		pi      float32
	}{false, 99, 90.2}

    e3 = e5

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println(e4)
	fmt.Println(e5)
}
