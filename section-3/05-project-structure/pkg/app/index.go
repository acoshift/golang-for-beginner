package app

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
