package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Config contains program configuration options.
var Config struct {
	Silent   bool
	DestFile string
}

func init() {
	// Let the flag package handle the options; -o for output and -s for silent
	flag.StringVar(&Config.DestFile, "o", "", "output file")
	flag.BoolVar(&Config.Silent, "s", false, "silent (do not output to stdout)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: ./io3 [options] <url>")
		os.Exit(2)
	}
}

func main() {
	// r here is a response, and r.Body is an io.Reader
	resp, err := http.Get(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	// A slice of io.Writers we will write the file to.
	var writers []io.Writer

	// If we aren't in Silent mode, lets add Stdout to our writers.
	if !Config.Silent {
		writers = append(writers, os.Stdout)
	}

	// If DestFile was provided, lets try to create it and to the writers.
	if len(Config.DestFile) > 0 {
		file, err := os.Create(Config.DestFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		writers = append(writers, file)
		defer file.Close()
	}

	// MultiWriter(io.Writer...) returns a single writer which multipleexes
	// its writes across all of the writers we pass in.
	dest := io.MultiWriter(writers...)

	// Write to dest the same way as before, copying from the Body.
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}
