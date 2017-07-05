package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(index))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Gopher.")
}
