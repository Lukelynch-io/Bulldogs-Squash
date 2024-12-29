package auth_test

import (
	"log"
	"testing"
	"webservice/app/auth"
	"webservice/app/blog/blog_claims"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

const username = "username"
const password = "password"

var testUser auth.User = auth.NewUser(username, password)

func TestCreateUser(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryAuthRepository)
	// Act
	newUserId, createUserError := auth.CreateUser(repo, username, password)
	if createUserError != nil {
		log.Fatal(createUserError)
	}
	actual, getUserError := repo.GetUserByUserId(newUserId)
	if getUserError != nil {
		log.Fatal(getUserError)
	}
	// Assert
	assert.IsEqual(actual.UserId, newUserId)
}

func TestGetCreatedUser(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryAuthRepository)
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		log.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := auth.GetUserByUserId(repo, testUser.UserId)
	if getUserError != nil {
		log.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestFailToCreateUserAlreadyExists(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryAuthRepository)
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		log.Fatal("Failed to create initial user")
	}
	// Act
	_, expectedCreateUserError := auth.CreateUser(repo, username, password)
	if expectedCreateUserError == nil {
		log.Fatal("Duplicate user was created")
	}

}

func TestGrantUserWithCreateClaim(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryAuthRepository)
	repo.CreateUser(testUser)
	// Act
	resultFlag, _ := auth.GrantUserClaim(repo, testUser.UserId, blog_claims.CREATE_BLOG)
	assert.Equal(t, resultFlag, true)
}
