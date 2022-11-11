package events

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Repository interface {
	GetEvents(ID string) (Events, error)
	// AddEvent(Event) (string, error)
}

type Service interface {
	GetEvents(ID string) (Events, error)
	// AddEvent(Event) (string, error)

}

type service struct {
	// r Repository
	event Event
	DB   firestore.Client
	ctx  context.Context
}

type Event struct {
	ID          string `json:"id"`
	Name		string `json:"name"`
	Location    string `json:"location"`
	Date		string `json:"date"`
	Time 	    string `json:"time"`
	TripId		string `json:"tripId"`
}

type Events []Event

func NewService(DB firestore.Client, ctx context.Context) *service {
	return &service{Event{}, DB, ctx}
}

func (s *service) GetEvents(ID string) (Events, error) {
	listOfAllEvents := Events{}

	resp, err := s.DB.Collection("events").Where("tripID", "==", ID).Documents(s.ctx).GetAll()
	if err != nil {
		return listOfAllEvents, err
	}

	fmt.Println(resp)

	for _, doc := range resp {
		var event Event
		fmt.Println(doc.Data())
		data := doc.Data()
		event = Event{
			ID:        doc.Ref.ID,
			Name:      data["name"].(string),
			Location:  data["location"].(string),
			Date:      data["date"].(string),
			Time:      data["time"].(string),
			TripId:    data["tripID"].(string),	
		}
		listOfAllEvents = append(listOfAllEvents, event)
	}

	return listOfAllEvents, nil
}