package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/trips"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
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

func updateTripHandler(tripsService trips.Service, usersService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tripId := mux.Vars(r)["TripId"]
		w.Header().Set("Access-Control-Allow-Origin", "*")

		type ShareUser struct {
			Email string `json:"email"`
			Editor bool `json:"editor"`
		}
		var shareUser ShareUser

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error reading body " + err.Error())
			return
		}

		err = json.Unmarshal(bodyBytes, &shareUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error unmarshalling body " + err.Error())
			return
		}

		userId, err := usersService.GetUserByEmail(shareUser.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error getting user by email " + err.Error())
			return
		}
		if userId.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("User not found")
			return
		}

		var updatedTrip trips.Trip
		updatedTrip.ID = tripId
		if shareUser.Editor {
			updatedTrip.Editors = append(updatedTrip.Editors, userId.ID)
		} else {
			updatedTrip.Viewers = append(updatedTrip.Viewers, userId.ID)
		}

		// fmt.Println("entering updateTripHandler")
		clearedTrip, err := tripsService.UpdateTrip(updatedTrip)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(clearedTrip)
	}
}