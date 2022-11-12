package handlers

import (
	"encoding/json"
	"net/http"
)

func healthcheck() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Welcome to the Trippin App. Everything looks good!")
	}
}