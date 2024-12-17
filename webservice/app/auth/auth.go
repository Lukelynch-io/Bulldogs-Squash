package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewToken(secretKey []byte, claims []Claim, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"claims":   claims,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	return token.SignedString(secretKey)
}
