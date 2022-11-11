package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolbyrogers/Trippin/backend/pkg/handlers"
)

func main() {
	fmt.Println("Starting Server on port 8080")
	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}