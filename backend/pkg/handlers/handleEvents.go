package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/events"
)

func getEventsHandler(eventsService events.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := mux.Vars(r)["TripId"]
		fmt.Println(ID)
		events, error := eventsService.GetEvents(ID)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode("Hit getUsersHandler")
		json.NewEncoder(w).Encode(events)
	}
}