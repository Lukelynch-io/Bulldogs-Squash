package app_test

import (
	"testing"
	"webservice/app"
	"webservice/app/blog/blog_claims"
	"webservice/domain"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestGrantUserWithCreateClaim(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	var testUser domain.User = domain.NewUser(username, password)
	repo := new(infra.MemoryAuthRepository)
	repo.CreateUser(testUser)
	// Act
	resultFlag, _ := app.UpdateUserClaims(repo, testUser.UserId, blog_claims.CREATE_BLOG)
	assert.Equal(t, resultFlag, true)
}
