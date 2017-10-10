package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
	// Create a value of type user.
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	fmt.Println(u)
	fmt.Println(&u)
}
