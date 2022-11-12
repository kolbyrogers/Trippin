package photos

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Service interface {
	GetPhotosByEventId(ID string) (Photos, error)
	GetPhotosByTripId(ID string) (Photos, error)
	AddPhoto(Photo) (error)
}

type service struct {
	photo Photo
	DB   firestore.Client
	ctx context.Context
}

type Photo struct {
	ID          string `json:"id"`
	TripId		string `json:"tripId"`
	EventId		string `json:"eventId"`
	Url			string `json:"photoUrl"`
}

type Photos []Photo

func NewService(DB firestore.Client, ctx context.Context) *service {
	return &service{Photo{}, DB, ctx}
}

func (s *service) GetPhotosByEventId(ID string) (Photos, error) {
	listOfAllPhotos := Photos{}

	resp, err := s.DB.Collection("photos").Where("eventId", "==", ID).Documents(s.ctx).GetAll()
	if err != nil {
		return listOfAllPhotos, err
	}

	for _, doc := range resp {
		var photo Photo
		data := doc.Data()
		photo = Photo{
			ID:        doc.Ref.ID,
			TripId:    data["tripId"].(string),
			EventId:   data["eventId"].(string),
			Url:       data["photoUrl"].(string),
		}
		listOfAllPhotos = append(listOfAllPhotos, photo)
	}

	return listOfAllPhotos, nil
}

func (s *service) GetPhotosByTripId(ID string) (Photos, error) {
	listOfAllPhotos := Photos{}

	resp, err := s.DB.Collection("photos").Where("tripId", "==", ID).Documents(s.ctx).GetAll()
	if err != nil {
		return Photos{}, err
	}

	for _, doc := range resp {
		var photo Photo
		data := doc.Data()
		photo = Photo{
			ID:        doc.Ref.ID,
			TripId:    data["tripId"].(string),
			EventId:   data["eventId"].(string),
			Url:       data["photoUrl"].(string),
		}
		listOfAllPhotos = append(listOfAllPhotos, photo)
	}

	return listOfAllPhotos, nil
}

func (s *service) AddPhoto(photo Photo) (error) {
	_, _, err := s.DB.Collection("photos").Add(s.ctx, map[string]interface{}{
		"tripId": photo.TripId,
		"eventId": photo.EventId,
		"photoUrl": photo.Url,
	})
	if err != nil {
		return err
	}

	return  nil
}