package utils

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func HashPassword(password string) (string, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), nil
}
