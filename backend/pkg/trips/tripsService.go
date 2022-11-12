package trips

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

type Repository interface {
	GetTrips(ID string) ([]Trip, error)
	AddTrip(Trip) (string, error)
	UpdateTrip(Trip) (Trip, error)
}

type Service interface {
	GetTrips(ID string) ([]Trip, error)
	AddTrip(Trip) (string, error)
	UpdateTrip(Trip) (Trip, error)

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
	Editors	    []interface{} `json:"editors"`
	Viewers	    []interface{} `json:"viewers"`
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
	
	// add a new trip to the trips collection
	for _, doc := range resp {
		var trip Trip
		trip.ID = doc.Ref.ID
		trip.Location = doc.Data()["location"].(string)
		trip.StartDate = doc.Data()["startDate"].(string)
		trip.EndDate = doc.Data()["endDate"].(string)
		for _, editor := range doc.Data()["editors"].([]interface{}) {
			trip.Editors = append(trip.Editors, editor.(string))
		}
		for _, viewer := range doc.Data()["viewers"].([]interface{}) {
			trip.Viewers = append(trip.Viewers, viewer.(string))
		}
		trip.ImageURL = doc.Data()["imageURL"].(string)
		listOfAllTrips = append(listOfAllTrips, trip)
	}
	for _, doc := range resp2 {
		var trip Trip
		trip.ID = doc.Ref.ID
		trip.Location = doc.Data()["location"].(string)
		trip.StartDate = doc.Data()["startDate"].(string)
		trip.EndDate = doc.Data()["endDate"].(string)
		for _, editor := range doc.Data()["editors"].([]interface{}) {
			trip.Editors = append(trip.Editors, editor.(string))
		}
		for _, viewer := range doc.Data()["viewers"].([]interface{}) {
			trip.Viewers = append(trip.Viewers, viewer.(string))
		}
		trip.ImageURL = doc.Data()["imageURL"].(string)
		listOfAllTrips = append(listOfAllTrips, trip)
	}
	
	return listOfAllTrips, nil
}

func (s *service) GetOneTrip(ID string) (Trip, error) {
	resp, err := s.DB.Collection("trips").Doc(ID).Get(s.ctx)
	if err != nil {
		fmt.Println("error getting trip" + err.Error())
		return Trip{}, err
	}
	// fmt.Println(resp.Data()["editors"])
	data := resp.Data()
	trip := Trip{
		ID: ID,
		Location: data["location"].(string),
		StartDate: data["startDate"].(string),
		EndDate: data["endDate"].(string),
		Editors: data["editors"].([]interface{}),
		Viewers: data["viewers"].([]interface{}),
		ImageURL: data["imageURL"].(string),
	}
	return trip, nil
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

func (s *service) UpdateTrip(passedTrip Trip) (Trip, error) {
	checkTrip, err := s.GetOneTrip(passedTrip.ID)
	if err != nil {
		fmt.Println("error on update trip with getting one trip" + err.Error())
		return Trip{}, err
	}

	if passedTrip.Location != "" {
		checkTrip.Location = passedTrip.Location
	}
	if passedTrip.StartDate != "" {
		checkTrip.StartDate = passedTrip.StartDate
	}
	if passedTrip.EndDate != "" {
		checkTrip.EndDate = passedTrip.EndDate
	}
	if passedTrip.Editors != nil {
		// check if the viewer is already in the list
		for _, newEditor := range passedTrip.Editors {
			shouldAdd := true
			for _, editor := range checkTrip.Editors {
				if newEditor == editor {
					shouldAdd = false
				}
			}
			if shouldAdd {
				checkTrip.Editors = append(checkTrip.Editors, newEditor)
			}
		}
	}
	if passedTrip.Viewers != nil {
		// check if the viewer is already in the list
		for _, newViewer := range passedTrip.Viewers {
			shouldAdd := true
			for _, viewer := range  checkTrip.Viewers {
				if viewer == newViewer {
					shouldAdd = false
				}
			}
			if shouldAdd {
				checkTrip.Viewers = append(checkTrip.Viewers, passedTrip.Viewers...)
			}
		}
	}
	if passedTrip.ImageURL != "" {
		checkTrip.ImageURL = passedTrip.ImageURL
	}

	_, err = s.DB.Collection("trips").Doc(checkTrip.ID).Set(s.ctx, map[string]interface{}{
		"location": checkTrip.Location,
		"startDate": checkTrip.StartDate,
		"endDate": checkTrip.EndDate,
		"editors": checkTrip.Editors,
		"viewers": checkTrip.Viewers,
		"imageURL": checkTrip.ImageURL,
	})
	if err != nil {
		fmt.Println(err)
		return Trip{}, err
	}

	changedTrip, err := s.GetOneTrip(checkTrip.ID)

	return changedTrip, nil
}

// 101 88 245