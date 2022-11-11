package handlers

import "github.com/gorilla/mux"

func InitializeHandlers() *mux.Router { // add this in a second r reading.Service, w writing.Service

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck()) //.Methods("GET")

	return router
}