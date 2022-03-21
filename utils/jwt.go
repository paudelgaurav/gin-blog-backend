package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateToken -> generates token
func GenerateToken(userid uint) string {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
		"iat":    time.Now().Unix(),
		"userID": userid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("abcdef"))
	return t

}
