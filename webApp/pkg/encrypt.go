package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(b), err
}

func ComparePass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
