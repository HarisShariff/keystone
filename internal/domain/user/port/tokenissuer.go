package port

// TokenIssuer defines methods for generating tokens
type TokenIssuer interface {
	GenerateToken(userID, role string) (string, error)
}
