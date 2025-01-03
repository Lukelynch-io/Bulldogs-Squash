package infra

import (
	"errors"
	"webservice/domain"
)

type MemoryAuthRepository struct {
	users  []domain.User
	tokens map[domain.UserId][]domain.TokenString
}

func NewMemoryAuthRepository() MemoryAuthRepository {
	return MemoryAuthRepository{
		users:  make([]domain.User, 0),
		tokens: make(map[domain.UserId][]domain.TokenString),
	}
}

func (repo *MemoryAuthRepository) GetUserByUserId(userId domain.UserId) (*domain.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			return &repo.users[i], nil
		}
	}
	return nil, errors.New("Could not find user")
}

func (repo *MemoryAuthRepository) GetUserByUsername(username string) (*domain.User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (repo *MemoryAuthRepository) CreateUser(newUser domain.User) (*domain.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Username == newUser.Username {
			return nil, errors.New("Username already taken. Please try another")
		}
	}
	repo.users = append(repo.users, newUser)

	return &newUser, nil
}

func (repo *MemoryAuthRepository) updateUser(userId domain.UserId, updatedUser domain.User) (bool, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			repo.users[i] = updatedUser
			return true, nil
		}
	}
	return false, errors.New("User not found")
}

func (repo *MemoryAuthRepository) GrantUserClaim(userId domain.UserId, newClaim domain.Claim) (bool, error) {
	foundUser, getUserError := repo.GetUserByUserId(userId)
	if getUserError != nil {
		return false, getUserError
	}
	foundUser.Claims[newClaim] = newClaim
	return true, nil
}

func (repo *MemoryAuthRepository) StoreToken(userId domain.UserId, tokenId domain.TokenString) error {
	repo.tokens[userId] = append(repo.tokens[userId], tokenId)
	return nil
}

func (repo *MemoryAuthRepository) RevokeToken(userId domain.UserId) error {
	delete(repo.tokens, userId)
	return nil
}
