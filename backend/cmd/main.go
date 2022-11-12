package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"github.com/kolbyrogers/Trippin/backend/pkg/events"
	"github.com/kolbyrogers/Trippin/backend/pkg/handlers"
	"github.com/kolbyrogers/Trippin/backend/pkg/photos"
	"github.com/kolbyrogers/Trippin/backend/pkg/trips"
	"github.com/kolbyrogers/Trippin/backend/pkg/users"
)

func main() {

	ctx := context.Background()
	DB := hookUpToDB(ctx)
	defer DB.Close()


	usersService := users.NewService(*DB, ctx)
	tripsService := trips.NewService(*DB, ctx)
	eventsService := events.NewService(*DB, ctx)
	photosService := photos.NewService(*DB, ctx)

	services := handlers.Services{
		UsersService: usersService,
		TripsService: tripsService,
		EventsService: eventsService,
		PhotosService: photosService,
	}

	fmt.Println("Starting Server on port 8080")
	router := handlers.InitializeHandlers(services)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hookUpToDB(ctx context.Context) (*firestore.Client) {
	// Sets your Google Cloud Platform project ID.
	projectID := "trippin-33f25"

	// Use a service account
	sa := option.WithCredentialsFile("./firestore-credentials.json")
	DB, err := firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return DB
}