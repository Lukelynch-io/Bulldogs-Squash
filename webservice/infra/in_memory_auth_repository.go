package infra

import (
	"errors"
	"webservice/app/auth"
	"webservice/app/auth/claim"
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

func (repo *MemoryAuthRepository) GetUserByUsername(username string) (auth.User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return user, nil
		}
	}
	return auth.User{}, errors.New("User not found")
}

func (repo *MemoryAuthRepository) CreateUser(newUser auth.User) (bool, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Username == newUser.Username {
			return false, errors.New("Username already taken. Please try another")
		}
	}
	repo.users = append(repo.users, newUser)

	return true, nil
}

func (repo *MemoryAuthRepository) updateUser(userId auth.UserId, updatedUser auth.User) (bool, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			repo.users[i] = updatedUser
			return true, nil
		}
	}
	return false, errors.New("User not found")
}

func (repo *MemoryAuthRepository) GrantUserClaim(userId auth.UserId, newClaim claim.Claim) (bool, error) {
	foundUser, getUserError := repo.GetUserByUserId(userId)
	if getUserError != nil {
		return false, getUserError
	}
	foundUser.Claims[newClaim] = newClaim
	return true, nil
}
