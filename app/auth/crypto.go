package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func Crypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}
