package app

import "webservice/domain"

func CreateUser(repo domain.IAuthRepo, username string, passwordHash string) (*domain.User, error) {
	newUser := domain.NewUser(username, passwordHash)
	_, createUserError := repo.CreateUser(newUser)
	if createUserError != nil {
		return nil, createUserError
	}
	return &newUser, nil
}
