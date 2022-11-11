package trips

import (
	"context"
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
)

type Repository interface {
	GetTrips(ID string) ([]Trip, error)
	AddTrip(Trip) (string, error)
}

type Service interface {
	GetTrips(ID string) ([]Trip, error)
	AddTrip(Trip) (string, error)

}

type service struct {
	// r Repository
	trip Trip
	DB   firestore.Client
	ctx  context.Context
}

type Trip struct {
	ID          string `json:"id"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Editors	    []string `json:"editors"`
	Viewers	    []string `json:"viewers"`
	ImageURL	string `json:"imageURL"`
}

func NewService(DB firestore.Client, ctx context.Context) *service {
	return &service{Trip{}, DB, ctx}
}

func (s *service) GetTrips(ID string) ([]Trip, error) {
	listOfAllTrips := []Trip{}
	resp, err := s.DB.Collection("trips").Where("editors", "array-contains", ID).Documents(s.ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		return listOfAllTrips, err
	}
	resp2, err := s.DB.Collection("trips").Where("viewers", "array-contains", ID).Documents(s.ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		return listOfAllTrips, err
	}
	// combine the two slices

	fmt.Print(resp)
	data := resp[0].Data()
	fmt.Print(data)
	// add a new trip to the trips collection


	for _, doc := range resp {
		fmt.Println(doc.Data())
		var trip Trip
		trip.ID = doc.Ref.ID
		trip.Location = doc.Data()["location"].(string)
		trip.StartDate = doc.Data()["startDate"].(time.Time).String()
		trip.EndDate = doc.Data()["endDate"].(time.Time).String()
		for _, editor := range doc.Data()["editors"].([]interface{}) {
			trip.Editors = append(trip.Editors, editor.(string))
		}
		for _, viewer := range doc.Data()["viewers"].([]interface{}) {
			trip.Viewers = append(trip.Viewers, viewer.(string))
		}
		trip.ImageURL = doc.Data()["imageURL"].(string)
		fmt.Println(trip.Location)
		listOfAllTrips = append(listOfAllTrips, trip)
	}
	for _, doc := range resp2 {
		fmt.Println(doc.Data())
		var trip Trip
		trip.ID = doc.Ref.ID
		trip.Location = doc.Data()["location"].(string)
		trip.StartDate = doc.Data()["startDate"].(time.Time).String()
		trip.EndDate = doc.Data()["endDate"].(time.Time).String()
		for _, editor := range doc.Data()["editors"].([]interface{}) {
			trip.Editors = append(trip.Editors, editor.(string))
		}
		for _, viewer := range doc.Data()["viewers"].([]interface{}) {
			trip.Viewers = append(trip.Viewers, viewer.(string))
		}
		trip.ImageURL = doc.Data()["imageURL"].(string)
		fmt.Println(trip.Location)
		listOfAllTrips = append(listOfAllTrips, trip)
	}
	
	return listOfAllTrips, nil
}

func (s *service) AddTrip(trip Trip) (string, error) {
	_, _, err := s.DB.Collection("trips").Add(s.ctx, map[string]interface{}{
		"location": trip.Location,
		"startDate": trip.StartDate,
		"endDate": trip.EndDate,
		"editors": trip.Editors,
		"viewers": trip.Viewers,
		"imageURL": trip.ImageURL,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return "successfully added trip", nil
}

// 101 88 245