package main

import (
	"log"
	"net/http"
)

func main() {
	h := m1(m2(m3(http.HandlerFunc(indexHandler))))
	http.ListenAndServe(":8080", h)
}

func m1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m1")
		h.ServeHTTP(w, r)
	})
}

func m2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m2")
		h.ServeHTTP(w, r)
	})
}

func m3(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m3")
		h.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
