package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {
	//users := make(map[string]user)
	//
	//users["Roy"] = user{"Rob", "Roy"}
	//users["Ford"] = user{"Henry", "Ford"}
	//users["Mouse"] = user{"Mickey", "Mouse"}
	//users["Jackson"] = user{"Michael", "Jackson"}

	users := map[string]user{
		"Roy":     user{"Rob", "Roy"},
		"Ford":    user{"Henry", "Ford"},
		"Mouse":   user{"Mickey", "Mouse"},
		"Jackson": user{"Michael", "Jackson"},
	}

	for k, v := range users {
		fmt.Println(k, v)
	}

	fmt.Println("--------------------------------------------------------------")

	for k, v := range users {
		fmt.Println(k, v)
	}

	delete(users, "Roy")

	k, found := users["Roy"]
	fmt.Println(k, found)
}
