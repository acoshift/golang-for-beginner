package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(indexHandler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher."))
}
