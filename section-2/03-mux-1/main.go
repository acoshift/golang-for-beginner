package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(mux))
}

func mux(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/about":
		aboutHandler(w, r)
	case "/login":
		loginHandler(w, r)
	default:
		indexHandler(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}
