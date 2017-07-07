package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

type indexData struct {
	Name string
	List []string
}

var t = template.Must(template.ParseFiles("index.tmpl"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := indexData{
		Name: "acoshift",
		List: []string{
			"Go",
			"C",
			"C++",
			"JavaScript",
		},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
