package users

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

type Repository interface {
	GetUser(ID string) (User, error)
	AddUser(User) (string, error)
	GetUserByEmail(email string) (User, error)
}

type Service interface {
	GetUser(ID string) (User, error)
	AddUser(User) (string, error)
	GetUserByEmail(email string) (User, error)
}

type service struct {
	// r Repository
	user User
	DB   firestore.Client
	ctx context.Context
}

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	
}

func NewService(DB firestore.Client, ctx context.Context) *service {
	return &service{User{}, DB, ctx}
}

func (s *service) GetUser(ID string) (User, error) {
	resp, err := s.DB.Collection("users").Doc(ID).Get(s.ctx)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	
	data := resp.Data()
	user := User{
		ID:        ID,
		DisplayName: data["displayName"].(string),
		Email:     data["email"].(string),
	}
	fmt.Println(data)

	
	return user, nil
}

func (s *service) AddUser(user User) (string, error) {

	// check for user exists

	_, _, err := s.DB.Collection("users").Add(s.ctx, map[string]interface{}{
		"displayName": user.DisplayName,
		"email": user.Email,
	})
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	return "successfully added a new user", nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	resp, err := s.DB.Collection("users").Where("email", "==", email).Documents(s.ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	
	// fmt.Print(resp)
	if(len(resp) > 0){
		data := resp[0].Data()
		user := User{
			ID:        resp[0].Ref.ID,
			DisplayName: data["displayName"].(string),
			Email:     data["email"].(string),
		}
	
	fmt.Println(data)
	return user, nil
	}
	return User{}, nil
}

// 101 88 245