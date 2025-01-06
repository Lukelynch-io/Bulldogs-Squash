package auth

func AuthenticateUser(authRepo UserDataStorage, secretKey []byte, username string, password string) (*TokenString, UserActionResponses) {
	foundUser, err := GetUserByUsername(authRepo, username)
	if err != nil {
		return nil, BadRequest
	}

	tokenString, err := GenerateNewToken(secretKey, foundUser.Claims.IntoArray(), *foundUser, nil)
	if err != nil {
		return nil, BadRequest
	}
	return tokenString, StatusOK
}

func RevokeUserToken(userStorage UserDataStorage, tokenStorage TokenStorage, username string) *UserActionResponses {
	var returnError UserActionResponses
	user, err := userStorage.GetUserByUsername(username)
	if err != nil {
		returnError = UserNotFoundError
		return &returnError
	}
	tokenStorage.RevokeToken(user.UserId)
	return nil
}
