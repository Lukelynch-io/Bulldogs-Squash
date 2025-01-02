package app

import "webservice/domain"

type RevokeUserTokenError string

const (
	UserNotFoundError RevokeUserTokenError = ("User not found")
)

func RevokeUserToken(repo domain.IAuthRepo, username string) *RevokeUserTokenError {
	var returnError RevokeUserTokenError
	user, err := repo.GetUserByUsername(username)
	if err != nil {
		returnError = UserNotFoundError
		return &returnError
	}
	repo.RevokeToken(user.UserId)
	return nil
}
