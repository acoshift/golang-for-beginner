package app

import "net/http"

func postView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post view"))
}
