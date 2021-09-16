package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
