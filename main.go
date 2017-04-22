package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// PORT port
	PORT = "127.0.0.1:8080"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sayHello)
	router.HandleFunc("/{name}", sayHelloTo)
	log.Println(http.ListenAndServe(PORT, router))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q, this is version 2", html.EscapeString(r.URL.Path))
}

func sayHelloTo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintln(w, "Hello,", name)
}
