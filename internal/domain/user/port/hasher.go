package port

// Hasher defines methods for hashing and verifying passwords
type Hasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}
