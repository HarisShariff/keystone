package user

import (
	"fmt"

	"github.com/HarisShariff/keystone/internal/domain/user/port"
)

// DummyTokenIssuer generates fake tokens for now
type DummyTokenIssuer struct{}

// NewDummyTokenIssuer creates a new DummyTokenIssuer
func NewDummyTokenIssuer() port.TokenIssuer {
	return &DummyTokenIssuer{}
}

func (t *DummyTokenIssuer) GenerateToken(userID, role string) (string, error) {
	// Just return a dummy token string
	return fmt.Sprintf("dummy-token-for-%s", userID), nil
}
