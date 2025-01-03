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
	testUser := createTestUser()
	repo := infra.NewMemoryAuthRepository()
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := app.GetUserByUsername(&repo, testUser.Username)
	if getUserError != nil {
		t.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestErrorIsReturnedWhenNoUserByUsername(t *testing.T) {
	// Arrange
	testUser := createTestUser()
	repo := infra.NewMemoryAuthRepository()
	// Act
	_, getUserError := app.GetUserByUsername(&repo, testUser.Username)
	if getUserError == nil {
		t.Fatal("User found when there should not have been")
	}
}

func TestGetCreatedUserById(t *testing.T) {
	// Arrange
	testUser := createTestUser()
	repo := infra.NewMemoryAuthRepository()
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := app.GetUserById(&repo, testUser.UserId)
	if getUserError != nil {
		t.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestErrorIsReturnedWhenNoUserByUserId(t *testing.T) {
	// Arrange
	testUser := createTestUser()
	repo := infra.NewMemoryAuthRepository()
	// Act
	_, getUserError := app.GetUserById(&repo, testUser.UserId)
	if getUserError == nil {
		t.Fatal("User found when there should not have been")
	}
}

func createTestUser() domain.User {
	const username = "username"
	const password = "password"
	return domain.NewUser(username, password, domain.Viewer)
}
