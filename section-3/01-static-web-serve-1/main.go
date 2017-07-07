package main

import "net/http"

func main() {
	http.HandleFunc("/", http.HandlerFunc(indexHandler))
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!doctype html>
		<html>
		<title>Static Web Server</title>
		<link href=/-/css/style.css rel=stylesheet>
		<p class=red>
			Static web server.
		</p>
	`))
}
