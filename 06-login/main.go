package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/kale/06-login/handler"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexGetHandler).Methods("GET")
	router.HandleFunc("/register", handler.RegisterGetHandler).Methods("GET")
	router.HandleFunc("/register", handler.RegisterPostHandler).Methods("POST")

	router.HandleFunc("/login", handler.LoginGetHandler).Methods("GET")
	router.HandleFunc("/login", handler.LoginPostHandler).Methods("POST")

	router.HandleFunc("/dashboard", handler.DashboardHandler).Methods("GET")

	http.ListenAndServe(":3030", router)
}
