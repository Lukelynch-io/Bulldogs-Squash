package app_test

import (
	"testing"
	"webservice/app"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestRevokeUserToken(t *testing.T) {
	// Arrange
	testAuthrepo := infra.NewMemoryAuthRepository()
	const username = "username"
	testAuthrepo.CreateUser(CreateTestUser())
	// Act
	revokeResult := app.RevokeUserToken(&testAuthrepo, username)
	// Assert
	assert.Equal(t, revokeResult, nil)
}

func TestRevokeUserTokenErrorsWhenUserIsntFound(t *testing.T) {
	// Arrange
	testAuthrepo := infra.NewMemoryAuthRepository()
	const username = "username"
	// Act
	revokeResult := app.RevokeUserToken(&testAuthrepo, username)
	// Assert
	assert.Equal(t, revokeResult, app.UserNotFoundError)
}
