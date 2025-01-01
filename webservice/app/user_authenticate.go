package app

import (
	"webservice/domain"
)

func AuthenticateUser(authRepo domain.IAuthRepo, secretKey []byte, username string, password string) (string, int) {
	foundUser, err := GetUserByUsername(authRepo, username)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	tokenString, err := domain.GenerateNewToken(secretKey, domain.IntoArray(foundUser.Claims), foundUser.Username, nil)
	if err != nil {
		return "", http.StatusBadRequest
	}
	return tokenString, 0
}
