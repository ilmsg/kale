package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", CounterHandler).Methods("GET")
	http.ListenAndServe(":3030", router)
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
