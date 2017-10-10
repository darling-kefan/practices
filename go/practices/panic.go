package main

/*
func main() {
	slice := make([]string, 2, 4)
	Example(slice, "hello", 10)
}

func Example(slice []string, str string, i int) {
	panic("Want stack trace")
}
*/

/*
import "fmt"

type trace struct{}

func main() {
	slice := make([]string, 2, 4)

	var t trace
	t.Example(slice, "hello", 10)
}

func (t *trace) Example(slice []string, str string, i int) {
	fmt.Printf("Receiver Address: %p\n", t)
	panic("Want stack trace")
}
*/

func main() {
	Example(true, false, true, 25)
}

func Example(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}
