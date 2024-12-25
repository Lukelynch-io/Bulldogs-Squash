package auth

import (
	"net/http"
	"time"
	"webservice/app/auth/claim"
	"webservice/app/blog/blog_claims"

	"github.com/golang-jwt/jwt"
)

func HandleUserAuth(secretKey []byte, username string, password string) (string, int) {
	// TODO: Do something with validation result
	_, err := ValidateUser(username, password)
	if err != nil {
		return "", http.StatusUnauthorized
	}
	claims := blog_claims.GetBlogClaims()

	tokenString, err := GenerateNewToken(secretKey, claims, "test_claim")
	if err != nil {
		return "", http.StatusBadRequest
	}
	return tokenString, 0
}

func ValidateUser(username string, password string) (bool, error) {
	return true, nil
}

func GenerateNewToken(secretKey []byte, claims []claim.Claim, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"claims":   claims,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	return token.SignedString(secretKey)
}
