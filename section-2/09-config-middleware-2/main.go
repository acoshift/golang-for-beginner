package main

import "net/http"

func main() {
	h := allowRoles("admin", "staff")(http.HandlerFunc(adminHandler))
	http.ListenAndServe(":8080", h)
}

type middleware func(http.Handler) http.Handler

func allowRoles(roles ...string) middleware {
	allow := make(map[string]struct{})
	for _, role := range roles {
		allow[role] = struct{}{}
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidded", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Admin and Staff."))
}
