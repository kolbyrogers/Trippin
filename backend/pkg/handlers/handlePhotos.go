package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/events"
	"github.com/kolbyrogers/Trippin/backend/pkg/photos"
	"github.com/kolbyrogers/Trippin/backend/pkg/trips"
)

func getPhotosHandlerByTrip(photosService photos.Service, tripsServices trips.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		vars := mux.Vars(r)
		tripId := vars["TripId"]
		photos, err := photosService.GetPhotosByTripId(tripId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(photos)
	}
}

func getPhotosHandlerByEvent(photosService photos.Service, eventsService events.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		vars := mux.Vars(r)
		eventId := vars["EventId"]

		event, err := eventsService.GetEvent(eventId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if event.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		photos, err := photosService.GetPhotosByEventId(eventId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(photos)
	}
}

func addPhotoHandler(photosService photos.Service, tripsService trips.Service, eventsService events.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		var photo photos.Photo
		err := json.NewDecoder(r.Body).Decode(&photo)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tripId, err := tripsService.GetTrip(photo.TripId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if tripId.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		eventId, err := eventsService.GetEvent(photo.EventId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if eventId.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = photosService.AddPhoto(photo)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}