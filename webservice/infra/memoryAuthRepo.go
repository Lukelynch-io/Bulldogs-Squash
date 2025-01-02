package infra

import (
	"errors"
	"webservice/domain"
)

type memoryAuthRepository struct {
	users  []domain.User
	tokens map[domain.UserId][]domain.TokenString
}

func NewMemoryAuthRepository() domain.IAuthRepo {
	return &memoryAuthRepository{
		users:  make([]domain.User, 0),
		tokens: make(map[domain.UserId][]domain.TokenString),
	}
}

func (repo *memoryAuthRepository) GetUserByUserId(userId domain.UserId) (*domain.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			return &repo.users[i], nil
		}
	}
	return nil, errors.New("Could not find user")
}

func (repo *memoryAuthRepository) GetUserByUsername(username string) (*domain.User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (repo *memoryAuthRepository) CreateUser(newUser domain.User) (*domain.User, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].Username == newUser.Username {
			return nil, errors.New("Username already taken. Please try another")
		}
	}
	repo.users = append(repo.users, newUser)

	return &newUser, nil
}

func (repo *memoryAuthRepository) updateUser(userId domain.UserId, updatedUser domain.User) (bool, error) {
	for i := 0; i < len(repo.users); i++ {
		if repo.users[i].UserId == userId {
			repo.users[i] = updatedUser
			return true, nil
		}
	}
	return false, errors.New("User not found")
}

func (repo *memoryAuthRepository) GrantUserClaim(userId domain.UserId, newClaim domain.Claim) (bool, error) {
	foundUser, getUserError := repo.GetUserByUserId(userId)
	if getUserError != nil {
		return false, getUserError
	}
	foundUser.Claims[newClaim] = newClaim
	return true, nil
}

func (repo *memoryAuthRepository) StoreToken(userId domain.UserId, tokenId domain.TokenString) error {
	repo.tokens[userId] = append(repo.tokens[userId], tokenId)
	return nil
}

func (repo *memoryAuthRepository) RevokeToken(userId domain.UserId) error {
	delete(repo.tokens, userId)
	return nil
}
