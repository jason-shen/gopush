package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash the user password
func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

// ComparePassword compare the user password with the given one to see if they match
func ComparePassword(attempt string, password string) error {
	bytePassword, byteHashedPassword := []byte(attempt), []byte(password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
