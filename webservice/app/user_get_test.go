package app_test

import (
	"testing"
	"webservice/app"
	"webservice/domain"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestGetCreatedUserByUsername(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	testUser := domain.NewUser(username, password, domain.Viewer)
	repo := new(infra.MemoryAuthRepository)
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := app.GetUserByUsername(repo, testUser.Username)
	if getUserError != nil {
		t.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestErrorIsReturnedWhenNoUserByUsername(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	testUser := domain.NewUser(username, password, domain.Viewer)
	repo := new(infra.MemoryAuthRepository)
	// Act
	_, getUserError := app.GetUserByUsername(repo, testUser.Username)
	if getUserError == nil {
		t.Fatal("User found when there should not have been")
	}
}

func TestGetCreatedUserById(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	testUser := domain.NewUser(username, password, domain.Viewer)
	repo := new(infra.MemoryAuthRepository)
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := app.GetUserById(repo, testUser.UserId)
	if getUserError != nil {
		t.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestErrorIsReturnedWhenNoUserByUserId(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	testUser := domain.NewUser(username, password, domain.Viewer)
	repo := new(infra.MemoryAuthRepository)
	// Act
	_, getUserError := app.GetUserById(repo, testUser.UserId)
	if getUserError == nil {
		t.Fatal("User found when there should not have been")
	}
}
