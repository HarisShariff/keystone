package auth

type UserRepository interface {
	CreateUser(username, passwordHash string) (string, error)
	GetUserByUsername(username string) (User, error)
}

type User struct {
	ID           string
	Username     string
	PasswordHash string
}

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}
