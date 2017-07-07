package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", http.HandlerFunc(indexHandler))
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(&noDirFS{http.Dir("static")})))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

type noDirFS struct {
	http.FileSystem
}

func (fs *noDirFS) Open(name string) (http.File, error) {
	f, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}
