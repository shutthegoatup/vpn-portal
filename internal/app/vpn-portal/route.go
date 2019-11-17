package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter()

	router.Use()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	router.HandleFunc("/issued", issuedHandler)
	router.HandleFunc("/", profilesHandler)
	router.HandleFunc("/profile/{profile}", viewHandler)
	router.HandleFunc("/profile/{profile}/issue", downloadHandler)

	return router
}
