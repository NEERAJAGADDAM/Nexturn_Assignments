package utils

import (
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

func ValidateJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})
	return err == nil && token.Valid
}
