package view

import (
	"html/template"
	"net/http"
)

var (
	tpIndex = template.Must(template.ParseFiles("templates/root.tmpl", "templates/index.tmpl")).Lookup("root")
)

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}

// Index renders index template
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
