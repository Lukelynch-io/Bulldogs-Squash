package infra

import (
	"errors"
	"webservice/pkg/auth"
)

type MemoryAuthRepository struct {
	users  []auth.User
	tokens map[auth.UserId][]auth.TokenString
}

func NewMemoryAuthRepository() MemoryAuthRepository {
	return MemoryAuthRepository{
		users:  make([]auth.User, 0),
		tokens: make(map[auth.UserId][]auth.TokenString),
	}
}

func (repo *MemoryAuthRepository) GetUserByUserId(userId auth.UserId) (*auth.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			return &repo.users[i], nil
		}
	}
	return nil, errors.New("Could not find user")
}

func (repo *MemoryAuthRepository) GetUserByUsername(username string) (*auth.User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (repo *MemoryAuthRepository) CreateUser(newUser auth.User) (*auth.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Username == newUser.Username {
			return nil, errors.New("Username already taken. Please try another")
		}
	}
	repo.users = append(repo.users, newUser)

	return &newUser, nil
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

func (repo *MemoryAuthRepository) GrantUserClaim(userId auth.UserId, newClaim auth.Claim) (bool, error) {
	foundUser, getUserError := repo.GetUserByUserId(userId)
	if getUserError != nil {
		return false, getUserError
	}
	foundUser.Claims[newClaim] = newClaim
	return true, nil
}

func (repo *MemoryAuthRepository) StoreToken(userId auth.UserId, tokenId auth.TokenString) error {
	repo.tokens[userId] = append(repo.tokens[userId], tokenId)
	return nil
}

func (repo *MemoryAuthRepository) RevokeToken(userId auth.UserId) error {
	delete(repo.tokens, userId)
	return nil
}
