package app

import "webservice/domain"

func GetUserByUsername(repo domain.IAuthRepo, username string) (*domain.User, error) {
	return repo.GetUserByUsername(username)
}
