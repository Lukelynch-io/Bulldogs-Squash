package app_test

import (
	"net/http"
	"testing"
	"webservice/app"
	"webservice/domain"
	"webservice/domain/blog_claims"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestGrantUserWithCreateClaim(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	var testUser domain.User = domain.NewUser(username, password, domain.Viewer)
	claim_map := make(map[domain.Claim]domain.Claim)
	claim_map[blog_claims.CREATE_BLOG] = blog_claims.CREATE_BLOG
	repo := infra.NewMemoryAuthRepository()
	repo.CreateUser(testUser)
	// Act
	result := app.UpdateUserClaims(repo, testUser.Username, claim_map)
	assert.Equal(t, result, http.StatusOK)
}
