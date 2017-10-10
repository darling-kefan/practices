package main

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
	mylog "mykit/log"
)

func main() {
	fmt.Printf("Ldate: %d\n", mylog.Ldate)
	fmt.Printf("Ltime: %d\n", mylog.Ltime)
	fmt.Printf("Lmicroseconds: %d\n", mylog.Lmicroseconds)
	fmt.Printf("Llongfile: %d\n", mylog.Llongfile)
	fmt.Printf("Lshortfile: %d\n", mylog.Lshortfile)
	fmt.Printf("LUTC: %d\n", mylog.LUTC)
	fmt.Printf("LstdFlags: %d\n", mylog.LstdFlags)
	fmt.Printf("Ldefault: %d\n", mylog.Ldefault)

	fmt.Printf("log.Ldate: %d\n", log.Ldate)
	fmt.Printf("log.Ltime: %d\n", log.Ltime)
	fmt.Printf("log.Lshortfile: %d\n", log.Lshortfile)

	fmt.Printf("log.Flags(): %d\n", log.Flags())

	logFile, err := filepath.Abs("log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(logFile)
	// fp, err := os.Open(logFile)
	fp, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	log.SetFlags(11)
	log.SetOutput(fp)
	log.Println("Hello World!")
}
