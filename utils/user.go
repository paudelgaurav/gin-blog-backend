package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) {
	bytePassword := []byte(*password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	*password = string(hashedPassword)
}
