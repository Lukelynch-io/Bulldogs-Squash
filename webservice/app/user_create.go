package app

import "webservice/domain"

func CreateUser(repo domain.IAuthRepo, username string, passwordHash string, role domain.UserRole) (*domain.User, error) {
	newUser := domain.NewUser(username, passwordHash, role)
	_, createUserError := repo.CreateUser(newUser)
	if createUserError != nil {
		return nil, createUserError
	}
	return &newUser, nil
}
