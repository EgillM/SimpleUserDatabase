package graph

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func comparePasswordAndHash(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
