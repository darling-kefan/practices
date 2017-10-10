package main

import (
	"fmt"
)

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")

	d.displayName()
	d.setAge(45)

	fmt.Println("What the Compiler is Doing:")
	data.displayName(d)
	(*data).setAge(&d, 45)

	// ***************************************************

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	f1()

	d.name = "Lisa"

	f1()

	// Declare a function variable for the method bound to the
	// d variable. The function variable will get the address of
	// d because the method is using a pointer receiver.
	f2 := d.setAge

	f2(45)

	d.name = "joan"

	f2(45)
}
