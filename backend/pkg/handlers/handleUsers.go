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
		ID := mux.Vars(r)["UserId"]
		fmt.Println(ID)
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
		var newUser users.User
		
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error reading request body " + err.Error())
			return
		}
		fmt.Println(string(bodyBytes[:]))

		// convert bodyBytes to a NewUser struct
		err = json.Unmarshal(bodyBytes, &newUser)

		// &userStruct.Email = string(bodyBytes["email"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Error unmarshalling request body " + err.Error())
			return
		}
		fmt.Println(newUser)

		_, err = usersService.AddUser(newUser)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
