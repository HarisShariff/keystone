package user

// Service defines business operations related to users
type Service interface {
	Register(username, password string) (*User, error)
	Authenticate(username, password string) (*User, error)
	GetByID(id string) (*User, error)
}
