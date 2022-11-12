package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/events"
	"github.com/kolbyrogers/Trippin/backend/pkg/trips"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)

func InitializeHandlers(usersService users.Service, tripsService trips.Service, eventsService events.Service) *mux.Router { // add this in a second r reading.Service, w writing.Service

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

	router.HandleFunc("/api/users", PreflightAddResourceHandler).Methods("OPTIONS")
	router.HandleFunc("/api/users/{UserId}", getUserHandler(usersService)).Methods("GET")
	router.HandleFunc("/api/users", addUserHandler(usersService)).Methods("POST")

	// Trips ----------------------------------------------------------------
	router.HandleFunc("/api/trips", PreflightAddResourceHandler).Methods("OPTIONS")
	router.HandleFunc("/api/trips/{UserId}", getTripHandler(tripsService)).Methods("GET")
	router.HandleFunc("/api/trips", addTripHandler(tripsService)).Methods("POST")
	router.HandleFunc("/api/trips/{TripId}", updateTripHandler(tripsService, usersService)).Methods("PUT")

	// Events ----------------------------------------------------------------
	router.HandleFunc("/api/events", PreflightAddResourceHandler).Methods("OPTIONS")
	router.HandleFunc("/api/events/{TripId}", getEventsHandler(eventsService)).Methods("GET")
	router.HandleFunc("/api/events", addEventHandler(eventsService)).Methods("POST")

	return router
}

func PreflightAddResourceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PreflightAddResourceHandler")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusNoContent)
}