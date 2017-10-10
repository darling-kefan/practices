package main

import (
	"fmt"
	"practices/toy"
)

func main() {
	// Use the New function from the toy package to create a value of
	// type toy.
	t := toy.New("shouqiang", 74)

	// Use the methods from the toy value to set some initialize
	// values.
	t.UpdateOnHand(10)
	t.UpdateSold(20)

	// Display each field separately from the toy value.
	fmt.Println(t.Name, t.Weight, t.OnHand(), t.Sold())
}
