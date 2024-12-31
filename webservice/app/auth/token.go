package auth

import (
	"fmt"
	"strings"
	"time"
	"webservice/app/auth/claim"

	"github.com/golang-jwt/jwt"
)

func GenerateNewToken(secretKey []byte, claims []claim.Claim, username string, creationTime *time.Time) (string, error) {

	if creationTime == nil {
		defaultCreationTime := time.Now().Add(time.Hour * 24)
		creationTime = &defaultCreationTime
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"claims":   claims,
			"exp":      creationTime.Unix(),
		})
	return token.SignedString(secretKey)
}

func ExtractBearerToken(authHeader string) (string, bool) {
	const bearerPrefix = "Bearer "
	return strings.CutPrefix(authHeader, bearerPrefix)
}

func ValidateToken(tokenString string, secretKey []byte) (*jwt.Token, bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	return token, token.Valid
}
