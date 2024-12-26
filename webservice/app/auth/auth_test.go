package auth_test

import (
	"log"
	"testing"
	"webservice/app/auth"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestCreateUser(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryAuthRepository)
	const username = "username"
	const password = "password"
	// Act
	auth.CreateUser(repo, username, password)
	// Arrange
	actual, err := auth.GetUserByUserId(repo, auth.UserId(username))
	// Assert
	if err != nil {
		log.Fatal(err)
	}
	assert.IsEqual(actual.UserId, auth.UserId(username))
}
