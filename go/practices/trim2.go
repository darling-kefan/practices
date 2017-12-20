package main

import (
    "fmt"
    "strings"
)

func main() {
    s := strings.TrimLeft(" hello world ", " ")
    fmt.Println(s)
}
