package port

import "github.com/HarisShariff/keystone/internal/domain/user"

// Repository defines methods for persisting and retrieving users
type Repository interface {
	Create(user *user.User) error
	GetByUsername(username string) (*user.User, error)
	GetByID(id string) (*user.User, error)
}
