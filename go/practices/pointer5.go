// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name string
	age  int
}

// Create a function that changes the value of one of the user fields.
func setName(u *user, name string) {

	fmt.Printf("address: %p, value: %p, point to: %v\n", &u, u, *u)

	// Use the pointer to change the value that the
	// pointer points to.
	u.name = name
}

// Create a function that changes the value of one of the user fields.
func setAge(u *user, age int) {

	// Use the pointer to change the value that the
	// pointer points to.
	u.age = age
}

func main() {

	// Create a variable of type user and initialize each field.
	u := user{"shouqiang", 30}

	// Display the value of the variable.
	fmt.Printf("address: %p, value: %v\n", &u, u)

	// Share the variable with the function you declared above.
	setName(&u, "chunhuan")
	setAge(&u, 33)

	// Display the value of the variable.
	fmt.Printf("address: %p, value: %v\n", &u, u)
}
