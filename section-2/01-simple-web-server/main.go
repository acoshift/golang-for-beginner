package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", &indexHandler{})
}

type indexHandler struct{}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gopher."))
}
