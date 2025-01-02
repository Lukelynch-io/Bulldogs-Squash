package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type customJwtClaims struct {
	Claims ClaimArray `json:"claims"`
	jwt.RegisteredClaims
}

func GenerateNewToken(secretKey []byte, claims []Claim, user User, expirationTime *time.Time) (string, error) {
	if expirationTime == nil {
		defaultCreationTime := time.Now().Add(time.Hour * 24).UTC()
		expirationTime = &defaultCreationTime
	}
	customClaims := customJwtClaims{
		claims,
		jwt.RegisteredClaims{
			Subject:   string(user.UserId),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(*expirationTime),
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	return token.SignedString(secretKey)
}

func ExtractBearerToken(authHeader string) (string, bool) {
	const bearerPrefix = "Bearer "
	return strings.CutPrefix(authHeader, bearerPrefix)
}

func ValidateToken(tokenString string, secretKey []byte) (*customJwtClaims, error) {
	var customClaims customJwtClaims
	token, err := jwt.ParseWithClaims(tokenString, &customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if token.Valid {
		return &customClaims, nil
	}
	return nil, err
}
