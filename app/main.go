package main

import (
	"fmt"
	"net/http"
)

func main() {
	defer fmt.Println("Server is shutting down...")
	fmt.Println("Starting server on port 8082...")

	http.HandleFunc("/", hello)

	http.ListenAndServe("0.0.0.0:8082", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}