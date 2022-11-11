package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)


func getUserHandler(usersService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		usersService.GetUser("1")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hit getUsersHandler")

	}
}

func addUserHandler(usersService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hit addUserHandler")
	}
}
