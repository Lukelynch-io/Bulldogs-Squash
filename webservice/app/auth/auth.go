package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewToken(secretKey []byte, claim string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": claim,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	return token.SignedString(secretKey)
}
