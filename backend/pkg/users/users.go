package users

import (
	"fmt"
	"net/http"
)

func GetUser(ID string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetUser")
	}
}

func AddUser(User) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AddUser")
	}
}