package main

import (
	"log"
	"net/http"
)

func main() {
	h := logger(http.HandlerFunc(indexHandler))
	http.ListenAndServe(":8080", h)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s\n", r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
