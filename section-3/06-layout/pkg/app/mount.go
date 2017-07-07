package app

import "net/http"

// Mount mounts handlers into mux
func Mount(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(index))

	adminMux := http.NewServeMux()
	adminMux.Handle("/", http.HandlerFunc(adminPosts))
	adminMux.Handle("/login", http.HandlerFunc(adminLogin))
	adminMux.Handle("/edit", http.HandlerFunc(adminPostEditor))

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
	mux.Handle("/post/", http.StripPrefix("/post", http.HandlerFunc(postView)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
