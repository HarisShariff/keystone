package user

import (
	"errors"
	"time"
)

// User represents a user in keystone
type User struct {
	ID        string // UUID
	Username  string // Username (or Email)
	Password  string // Hashed
	Role      string // User Role (admin, user, etc.)
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id, username, password, role string) (*User, error) {
	if username == "" || password == "" {
		return nil, errors.New("invalid user input")
	}
	now := time.Now()
	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
