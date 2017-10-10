package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	u := user{"Bill", "bill@email.com"}

	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
