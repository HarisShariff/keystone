package user

import (
	"errors"
	"sync"

	domain "github.com/HarisShariff/keystone/internal/domain/user"
	"github.com/HarisShariff/keystone/internal/domain/user/port"
)

// Memory Repository is an in-memory implementation of user.Repository. It is safe for concurrent use
type MemoryRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewMemoryRepository() port.Repository {
	return &MemoryRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *MemoryRepository) Create(user *domain.User) error {
	r.mu.Lock()         // Full Lock: Because we're modifying the map (write operation)
	defer r.mu.Unlock() // Unlock after done

	r.users[user.ID] = user
	return nil
}

func (r *MemoryRepository) GetByUsername(username string) (*domain.User, error) {
	r.mu.RLock()         // Read Lock: Safe for concurrent reads
	defer r.mu.RUnlock() // Unlock after done
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("User Not Found")
}

func (r *MemoryRepository) GetByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("User Not Found")
	}
	return user, nil
}
