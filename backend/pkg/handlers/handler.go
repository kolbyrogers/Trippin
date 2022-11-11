package handlers

import (
	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)

func InitializeHandlers(usersService users.Service) *mux.Router { // add this in a second r reading.Service, w writing.Service

	// tempuser := users.User{
	// 	ID:        "1",
	// 	FirstName: "Kolby",
	// 	LastName:  "Rogers",
	// 	Email:     "kolbyrogers1@gmail.com",
	// }

	// id := "1"

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck()) //.Methods("GET")

	// Users ----------------------------------------------------------------
	// router.HandleFunc("/users", getUserHandler(usersService)).Methods("GET")
	// router.HandleFunc("/users", addUserHandler(usersService)).Methods("POST")

	router.HandleFunc("/api/users/{UserId}", getUserHandler(usersService)).Methods("GET")
	router.HandleFunc("/api/users", addUserHandler(usersService)).Methods("POST")

	return router
}