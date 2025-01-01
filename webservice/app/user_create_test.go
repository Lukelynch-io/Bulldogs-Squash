package app_test

import (
	"testing"
	"webservice/app"
	"webservice/domain"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestCreateUser(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	repo := new(infra.MemoryAuthRepository)
	// Act
	newUser, createUserError := app.CreateUser(repo, username, password)
	if createUserError != nil {
		t.Fatal(createUserError)
	}
	actual, getUserError := repo.GetUserByUsername(newUser.Username)
	if getUserError != nil {
		t.Fatal(getUserError)
	}
	// Assert
	assert.Equal(t, actual, newUser)
}

func TestFailToCreateUserAlreadyExists(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	var testUser domain.User = domain.NewUser(username, password)
	repo := new(infra.MemoryAuthRepository)
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	_, expectedCreateUserError := app.CreateUser(repo, username, password)
	if expectedCreateUserError == nil {
		t.Fatal("Duplicate user was created")
	}

}
