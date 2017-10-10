package main

import (
	"fmt"
	"practices/users"
)

func main() {
	m := users.Manager{Title: "Hello"}

	m.ID = 1
	m.Name = "shouqiang"

	fmt.Println(m)
}
