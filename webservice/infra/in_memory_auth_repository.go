package infra

import (
	"errors"
	"webservice/app/auth"
)

type MemoryAuthRepository struct {
	users []auth.User
}

func (repo *MemoryAuthRepository) GetUserByUserId(userId auth.UserId) (auth.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			return repo.users[i], nil
		}
	}
	return auth.User{}, errors.New("Could not find user")
}

func (repo *MemoryAuthRepository) CreateUser(username string, passwordHash string) (auth.UserId, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Username == username {
			return auth.UserId(""), errors.New("Username already taken. Please try another")
		}
	}
	newUser := auth.User{
		UserId:       auth.UserId(username),
		Username:     username,
		PasswordHash: passwordHash,
	}
	repo.users = append(repo.users, newUser)

	return newUser.UserId, nil
}
