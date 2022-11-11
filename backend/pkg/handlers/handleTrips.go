package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/trips"
)


func getTripHandler(tripsService trips.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := mux.Vars(r)["UserId"]
		fmt.Println(ID)
		trips, error := tripsService.GetTrips(ID)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode("Hit getUsersHandler")
		json.NewEncoder(w).Encode(trips)
	}
}

func addTripHandler(tripsService trips.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newTrip trips.Trip
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error reading request body " + err.Error())
			return
		}
		fmt.Println(string(bodyBytes[:]))
		// convert bodyBytes to a NewUser struct
		err = json.Unmarshal(bodyBytes, &newTrip)
		// &userStruct.Email = string(bodyBytes["email"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error unmarshalling request body " + err.Error())
			return
		}
		fmt.Println(newTrip)

		createdTrip, err := tripsService.AddTrip(newTrip)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTrip)
	}
}