package app

import "webservice/domain"

func GetUserByUsername(repo domain.IAuthRepo, username string) (*domain.User, error) {
	return repo.GetUserByUsername(username)
}

func GetUserById(repo domain.IAuthRepo, userId domain.UserId) (*domain.User, error) {
	return repo.GetUserByUserId(userId)
}
