package auth

import (
	"fmt"
	"net/http"
	"webservice/app/auth/claim"
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

	tokenString, err := GenerateNewToken(secretKey, claim.ClaimMapIntoArray(foundUser.Claims), "test_claim", nil)
	if err != nil {
		return "", http.StatusBadRequest
	}
	return tokenString, 0
}
