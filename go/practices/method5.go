package main

import "fmt"

func event(message string) {
	fmt.Println(message)
}

type data struct {
	name string
	age  int
}

func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

func fireEvent1(f func(string)) {
	f("anonymous")
}

type handler func(string)

func fireEvent2(h handler) {
	h("handler")
}

func main() {
	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	// Use the fireEvent1 handler accepts any function
	// or method with the right signature.
	fireEvent1(event)
	fireEvent1(d.event)

	fireEvent2(event)
	fireEvent2(d.event)

	f1 := handler(event)
	f2 := handler(d.event)

	fireEvent1(f1)
	fireEvent1(f2)

	fireEvent2(f1)
	fireEvent2(f2)
}
