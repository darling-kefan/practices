package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Printf("User Name: %s\n", u.name)
}

func main() {
	u := user{"Bill"}

	entities := []printer{
		u,
		&u,
	}

	u.name = "Bill_CHG"

	for _, e := range entities {
		e.print()
	}
}
