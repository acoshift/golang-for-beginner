package main

import (
	"net/http"

	"github.com/acoshift/golang-for-beginner/section-3/06-layout/pkg/app"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(":8080", mux)
}
