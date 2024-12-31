package auth

import (
	"fmt"
	"strings"
	"time"
	"webservice/app/auth/claim"

	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	Username string        `json:"username"`
	Claims   []claim.Claim `json:"claims"`
	jwt.RegisteredClaims
}

func GenerateNewToken(secretKey []byte, claims []claim.Claim, username string, expirationTime *time.Time) (string, error) {

	if expirationTime == nil {
		defaultCreationTime := time.Now().Add(time.Hour * 24)
		expirationTime = &defaultCreationTime
	}
	customClaims := CustomJwtClaims{
		username,
		claims,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(*expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	return token.SignedString(secretKey)
}

func ExtractBearerToken(authHeader string) (string, bool) {
	const bearerPrefix = "Bearer "
	return strings.CutPrefix(authHeader, bearerPrefix)
}

func ValidateToken(tokenString string, secretKey []byte) (*jwt.Token, CustomJwtClaims, bool) {
	var customClaims CustomJwtClaims
	token, _ := jwt.ParseWithClaims(tokenString, &customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	return token, customClaims, token.Valid
}
