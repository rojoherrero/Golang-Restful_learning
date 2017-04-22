package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

const (
	// PORT port
	PORT = "127.0.0.1:8080"
)

func main() {
	log.Print("Running server on " + PORT)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(PORT, nil))
}
