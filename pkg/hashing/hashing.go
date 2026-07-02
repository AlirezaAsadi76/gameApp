package hashing

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hashing from a plaintext string

func HashPassword(password string) (string, error) {
	// GenerateFromPassword expects a byte slice and a cost factor
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a plaintext password with a hashed password

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
