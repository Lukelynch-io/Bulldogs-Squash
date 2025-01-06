package auth

import (
	"errors"
)

type InMemoryAuthRepository struct {
	users  []User
	tokens map[UserId][]TokenString
}

func NewInMemoryAuthRepository() InMemoryAuthRepository {
	return InMemoryAuthRepository{
		users:  make([]User, 0),
		tokens: make(map[UserId][]TokenString),
	}
}

func (repo *InMemoryAuthRepository) GetUserByUserId(userId UserId) (*User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			return &repo.users[i], nil
		}
	}
	return nil, errors.New("Could not find user")
}

func (repo *InMemoryAuthRepository) GetUserByUsername(username string) (*User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (repo *InMemoryAuthRepository) CreateUser(newUser User) (*User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Username == newUser.Username {
			return nil, errors.New("Username already taken. Please try another")
		}
	}
	repo.users = append(repo.users, newUser)

	return &newUser, nil
}

func (repo *InMemoryAuthRepository) updateUser(userId UserId, updatedUser User) (bool, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			repo.users[i] = updatedUser
			return true, nil
		}
	}
	return false, errors.New("User not found")
}

func (repo *InMemoryAuthRepository) GrantUserClaim(userId UserId, newClaim Claim) (bool, error) {
	foundUser, getUserError := repo.GetUserByUserId(userId)
	if getUserError != nil {
		return false, getUserError
	}
	foundUser.Claims[newClaim] = newClaim
	return true, nil
}

func (repo *InMemoryAuthRepository) StoreToken(userId UserId, tokenId TokenString) error {
	repo.tokens[userId] = append(repo.tokens[userId], tokenId)
	return nil
}

func (repo *InMemoryAuthRepository) RevokeToken(userId UserId) error {
	delete(repo.tokens, userId)
	return nil
}
