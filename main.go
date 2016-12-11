package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func levi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, Levi.")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/levi", levi)
	http.ListenAndServe(":8000", nil)
}
