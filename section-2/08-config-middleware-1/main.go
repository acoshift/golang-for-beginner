package main

import "net/http"

func main() {
	h := allowRole("admin")(http.HandlerFunc(adminHandler))
	http.ListenAndServe(":8080", h)
}

type middleware func(http.Handler) http.Handler

func allowRole(role string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if reqRole != role {
				http.Error(w, "Forbidded", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Admin."))
}
