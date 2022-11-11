package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolbyrogers/Trippin/backend/pkg/handlers"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)

func main() {

	usersService := users.NewService()

	fmt.Println("Starting Server on port 8080")
	router := handlers.InitializeHandlers(usersService)
	log.Fatal(http.ListenAndServe(":8080", router))
}