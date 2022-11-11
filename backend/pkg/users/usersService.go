package users

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

type Repository interface {
	GetUser(ID string) (User, error)
	// AddUser(User) (string, error)
}

type Service interface {
	GetUser(ID string) (User, error)
	// AddUser(User) (string, error)
}

type service struct {
	// r Repository
	user User
	DB   firestore.Client
	ctx context.Context
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	
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

	fmt.Println(data)
	// fmt.Println(resp)
	// body, err := ioutil.ReadAll(resp)

	// // unmarshal user into User struct
	// returnuser, err := Unmarshal(body, &User{})
	// if err != nil {
	// 	fmt.Println(err)
	// }


	// fmt.Println(user)
	
	return User{}, nil
}

// 101 88 245