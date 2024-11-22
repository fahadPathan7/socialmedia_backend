package batman

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Generate a salted hash for the input string
func Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// Compare string to generated hash
func Compare(hash string, s string) bool {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming) == nil
}

// ValidatePassword compares the password provided to that stored as a hash in the database
func ValidatePassword(ctx context.Context, password string, passwordHash string) error {

	if !Compare(passwordHash, password) {
		return fmt.Errorf("The password does not match")
	}

	return nil
}
