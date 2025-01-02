package app

import (
	"net/http"
	"webservice/domain"
)

func AuthenticateUser(authRepo domain.IAuthRepo, secretKey []byte, username string, password string) (*string, int) {
	foundUser, err := GetUserByUsername(authRepo, username)
	if err != nil {
		return nil, http.StatusUnauthorized
	}

	tokenString, err := domain.GenerateNewToken(secretKey, foundUser.Claims.IntoArray(), *foundUser, nil)
	if err != nil {
		return nil, http.StatusBadRequest
	}
	return &tokenString, http.StatusOK
}
