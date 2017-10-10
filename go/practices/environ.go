package main

import (
	"os"
	"fmt"
)

func main() {
	envs := os.Environ()
	for _, val := range envs {
		fmt.Println(val)
	}
}
