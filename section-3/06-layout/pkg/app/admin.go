package app

import "net/http"

func adminLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Login"))
}

func adminPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Posts"))
}

func adminPostEditor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Post Editor"))
}
