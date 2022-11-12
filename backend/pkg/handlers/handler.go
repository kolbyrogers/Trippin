package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/events"
	"github.com/kolbyrogers/Trippin/backend/pkg/photos"
	"github.com/kolbyrogers/Trippin/backend/pkg/trips"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)

type Services struct {
	UsersService users.Service
	TripsService trips.Service
	EventsService events.Service
	PhotosService photos.Service
}

func InitializeHandlers(services Services) *mux.Router { // add this in a second r reading.Service, w writing.Service

	usersService := services.UsersService
	tripsService := services.TripsService
	eventsService := services.EventsService
	photosService := services.PhotosService

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck()) //.Methods("GET")

	// Users ----------------------------------------------------------------
	router.HandleFunc("/api/users", PreflightAddResourceHandler).Methods("OPTIONS")
	router.HandleFunc("/api/users/{UserId}", getUserHandler(usersService)).Methods("GET")
	router.HandleFunc("/api/users", addUserHandler(usersService)).Methods("POST")

	// Trips ----------------------------------------------------------------
	router.HandleFunc("/api/trips", PreflightAddResourceHandler).Methods("OPTIONS")
	router.HandleFunc("/api/trips/users/{UserId}", getTripsHandler(tripsService)).Methods("GET")
	router.HandleFunc("/api/trips", addTripHandler(tripsService)).Methods("POST")
	router.HandleFunc("/api/trips/{TripId}", getTripHandler(tripsService)).Methods("GET")
	router.HandleFunc("/api/trips/{TripId}", updateTripHandler(tripsService, usersService)).Methods("PUT")
	router.HandleFunc("/api/trips/{TripId}", PreflightAddResourceHandler).Methods("OPTIONS")
	

	// Events ----------------------------------------------------------------
	router.HandleFunc("/api/events", PreflightAddResourceHandler).Methods("OPTIONS")
	router.HandleFunc("/api/events/trips/{TripId}", getEventsHandler(eventsService)).Methods("GET") // updated this one
	router.HandleFunc("/api/events/{EventId}", addEventHandler(eventsService)).Methods("GET")
	router.HandleFunc("/api/events", addEventHandler(eventsService)).Methods("POST") // want to change this endpoint

	// Photos ----------------------------------------------------------------
	router.HandleFunc("/api/trips/{TripId}/photos", getPhotosHandler(photosService)).Methods("GET")
	// router.HandleFunc("/api/photos", addPhotoHandler(photosService)).Methods("POST")

	return router
}

func PreflightAddResourceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PreflightAddResourceHandler")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, x-requested-with")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusNoContent)
	return
}