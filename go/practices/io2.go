package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./io2 <url>")
		os.Exit(2)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	// io.Copy(dst io.Writer, src io.Reader) (int64, err)
	n, err := io.Copy(os.Stdout, resp.Body)
	fmt.Println(n, err)
}
