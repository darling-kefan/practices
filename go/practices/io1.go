package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	n, err := b.Write([]byte("Hello "))
	fmt.Println(n, err)

	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)
}
