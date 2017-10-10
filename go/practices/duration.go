package main

import "fmt"

type duration int

func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func main() {
	d := duration(2)
	(&d).notify()
}
