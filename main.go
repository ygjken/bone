package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("web/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	files := []string{
		"web/templates/body.html",
		"web/templates/layout.html",
		"web/templates/navbar.html",
	}

	t = template.Must(t.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", nil)
}
