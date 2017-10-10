package main

// Sample program that adds a few more features. If -z is passed, we want
// any DestFile's to be gzipped. If -md5 is passed, we want print the md5sum
// of the data that's been transferred instead of the data itself.
import (
	"compress/gzip"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Config contains program configuration options.
var Config struct {
	Silent   bool
	DestFile string
	Gzip     bool
	Md5      bool
}

// init is called before main.
func init() {
	flag.StringVar(&Config.DestFile, "o", "", "output file")
	flag.BoolVar(&Config.Silent, "s", false, "silent (do not output to stdout)")
	flag.BoolVar(&Config.Gzip, "z", false, "gzip file output")
	flag.BoolVar(&Config.Md5, "md5", false, "stdout md5sum instead of body")
	flag.Parse()

	if len(flag.Args()) != 1 {
		// Determine the program name dynamically.
		prog := filepath.Base(os.Args[0])
		fmt.Printf("Usage: %s %s [options] <url>\n", prog, os.Args[0])
		os.Exit(2)
	}
}

func main() {
	// Capture the url from the arguments
	url := flag.Args()[0]

	// resp here is a response, and resp.Body is an io.Reader
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Our Md5 hash destination, which is an io.Writer that computes the
	// hash of whatever is written to it.
	hash := md5.New()
	var writers []io.Writer

	// If we aren't in Silent mode, lets add Stdout to our writers.
	if !Config.Silent {
		// If -md5 was passed, write to the hash instead of os.Stdout
		if Config.Md5 {
			writers = append(writers, hash)
		} else {
			writers = append(writers, os.Stdout)
		}
	}

	// If DestFile was provided, lets try to create it and add to the writers
	if len(Config.DestFile) > 0 {
		// By declaring a Writer here as a WriterClose, we're saying that we
		// don't care what the underlying implementation is at all, all we
		// require is something that can Write and Close; both os.File and
		// the gzip.Writer are WriteClosers.
		var writer io.WriteCloser
		writer, err := os.Create(Config.DestFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		// If we're in Gzip mode, wrap the writer in gzip
		if Config.Gzip {
			writer = gzip.NewWriter(writer)
		}

		writers = append(writers, writer)
	}

	// MultiWriter(io.Writer...) return a single writer which multipleexes
	// its writes across all of the writers we pass in.
	dest := io.MultiWriter(writers...)

	// Write to dest the same way as before, copying from the Body.
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// If we were in Md5 output mode, lets output the checksum and url.
	if Config.Md5 {
		fmt.Printf("%x %s\n", hash.Sum(nil), url)
	}
}
