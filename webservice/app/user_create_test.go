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
	var repo domain.IAuthRepo
	repo = infra.NewMemoryAuthRepository()
	// Act
	newUser, createUserError := app.CreateUser(repo, username, password, domain.Viewer)
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
	repo := infra.NewMemoryAuthRepository()
	_, createUserError := app.CreateUser(repo, username, password, domain.Viewer)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	_, expectedCreateUserError := app.CreateUser(repo, username, password, domain.Viewer)
	if expectedCreateUserError == nil {
		t.Fatal("Duplicate user was created")
	}

}
