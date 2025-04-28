package user

import (
	"github.com/HarisShariff/keystone/internal/domain"
	"github.com/HarisShariff/keystone/internal/domain/user/port"
	"github.com/google/uuid"
)

// Service implements domain.UserService interface
type Service struct {
	repo   port.Repository
	hasher port.Hasher
	issuer port.TokenIssuer
}

func NewService(repo port.Repository, hasher port.Hasher, issuer port.TokenIssuer) *Service {
	return &Service{
		repo:   repo,
		hasher: hasher,
		issuer: issuer,
	}
}

func (s *Service) Register(username, password string) (*domain.User, error) {
	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	userID := uuid.NewString()
	user, err := domain.NewUser(userID, username, hashedPassword, "user") // Default role to User
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Authenticate(username, password string) (*domain.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	// What happened to hashing?
	if err := s.hasher.Compare(user.Password, password); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByID(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}
