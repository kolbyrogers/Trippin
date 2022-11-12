package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)

func getUserHandler(usersService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		ID := mux.Vars(r)["UserId"]
		user, error := usersService.GetUser(ID)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode("Hit getUsersHandler")
		json.NewEncoder(w).Encode(user)

	}
}

func addUserHandler(usersService users.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		var newUser users.User

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error reading request body " + err.Error())
			return
		}

		// convert bodyBytes to a NewUser struct
		err = json.Unmarshal(bodyBytes, &newUser)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error unmarshalling request body " + err.Error())
			return
		}

		currentUser, _ := usersService.GetUserByEmail(newUser.Email)
		if currentUser.Email != "" {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode("User already exists")
			return
		}


		fmt.Println("attempting to create user with data: ", newUser)

		createdUser, err := usersService.AddUser(newUser)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("created User: " + createdUser)
	}
}