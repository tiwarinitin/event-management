package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
