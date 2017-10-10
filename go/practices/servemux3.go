package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

/*
func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
*/

func main() {
	mux := http.NewServeMux()

	// Convert the timeHandler function to HandlerFunc type
	// th := http.HandlerFunc(timeHandler)
	th := timeHandler(time.RFC3339)
	// Add th to mux
	mux.Handle("/time", th)

	log.Println("listening...")
	http.ListenAndServe(":3000", mux)
}
