package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Site struct {
	Title       string
	Description string
}

var site = &Site{
	Title:       "My WebSite",
	Description: "My WebSite Description",
}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	r.HandleFunc("/", indexHandler).Methods(http.MethodGet)

	http.ListenAndServe(":3070", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, site, []string{"index.html"})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{filepath.Join("templates", "layout.html")}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
