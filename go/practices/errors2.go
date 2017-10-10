package main

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest       = errors.New("Bad Request")
	ErrMovedPermanently = errors.New("Moved Permanently")
)

func main() {
	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return
		case ErrMovedPermanently:
			fmt.Println("The URL moved, check it again")
			return
		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Life is good")
}

// webCall performs a web operation.
func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}
	return ErrMovedPermanently
}
