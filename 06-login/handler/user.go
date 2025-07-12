package handler

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/ilmsg/kale/06-login/model"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	site := &model.Site{
		Title:       "My WebSite",
		Description: "My WebSite Description",
		Author:      "Eak Netpanya",
	}

	Render(w, site, []string{"index.html"})
}

func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"register.html"})
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password cannot be empty", http.StatusBadRequest)
		return
	}

	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"login.html"})
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"dashboard.html"})
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	counter := session.Values["counter"]
	if counter == nil {
		counter = 0
	}
	currentCounter := counter.(int) + 1
	session.Values["counter"] = currentCounter
	session.Save(r, w)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"counter": currentCounter,
	})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{filepath.Join("templates", "layout.html")}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
