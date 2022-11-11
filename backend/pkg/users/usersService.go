package users

// package firebase




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
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	
}

func NewService() *service {
	return &service{User{}}
}

func (s *service) GetUser(ID string) (User, error) {
	return User{}, nil
}

// 101 88 245