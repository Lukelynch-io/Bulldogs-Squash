package auth

import (
	"fmt"
	"net/http"
	"time"
	"webservice/app/auth/claim"

	"github.com/golang-jwt/jwt"
)

type IAuthRepo interface {
	GetUserByUserId(UserId) (User, error)
	GetUserByUsername(string) (User, error)
	CreateUser(User) (bool, error)
	GrantUserClaim(UserId, claim.Claim) (bool, error)
}

func CreateUser(repo IAuthRepo, username string, passwordHash string) (UserId, error) {
	newUser := NewUser(username, passwordHash)
	_, createUserError := repo.CreateUser(newUser)
	if createUserError != nil {
		return UserId(""), createUserError
	}
	return newUser.UserId, nil
}

func GetUserByUserId(repo IAuthRepo, userId UserId) (User, error) {
	return repo.GetUserByUserId(userId)
}

func GrantUserClaim(repo IAuthRepo, userId UserId, newClaim claim.Claim) (bool, error) {
	// TODO: add executing user verification
	return repo.GrantUserClaim(userId, newClaim)
}

func HandleUserAuth(authRepo IAuthRepo, secretKey []byte, username string, password string) (string, int) {
	// TODO: Do something with validation result
	foundUser, err := authRepo.GetUserByUsername(username)
	fmt.Printf("Found user ID: %s", foundUser.UserId)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	tokenString, err := GenerateNewToken(secretKey, claim.ClaimMapIntoArray(foundUser.Claims), "test_claim")
	if err != nil {
		return "", http.StatusBadRequest
	}
	return tokenString, 0
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
