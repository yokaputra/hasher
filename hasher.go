package hasher

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Hasher represents a hasher instance.
type Hasher struct{}

// IHasher is the interface for hashing operations.
type IHasher interface {
	Hash(b []byte, cost int) ([]byte, error)
	Verify(plain, hashed string) bool
	Generate(cost int) (string, error)
}

// NewHasher creates a new Hasher instance.
func NewHasher() IHasher {
	return &Hasher{}
}

const (
	// DefaultCost is the default cost for hashing.
	DefaultCost = 12
)

// Hash generates a bcrypt hash from the given byte slice.
func (*Hasher) Hash(b []byte, cost int) ([]byte, error) {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		return nil, errors.New("invalid cost value")
	}
	return bcrypt.GenerateFromPassword(b, cost)
}

// Verify checks if a plain text matches a hashed password.
func (*Hasher) Verify(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

// Generate generates a random string of the specified length.
func (*Hasher) Generate(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid length")
	}

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&()[]{}<>?")
	res := make([]rune, length)

	for i := range res {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		res[i] = letters[int(b[0])%len(letters)]
	}

	return string(res), nil
}
