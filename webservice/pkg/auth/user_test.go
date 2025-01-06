package auth_test

import (
	"testing"
	"webservice/infra"
	"webservice/pkg/auth"

	"github.com/go-playground/assert/v2"
)

func TestCreateUser(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	repo := infra.NewMemoryAuthRepository()
	// Act
	newUser, createUserError := auth.CreateUser(&repo, username, password, auth.Viewer)
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
	_, createUserError := auth.CreateUser(&repo, username, password, auth.Viewer)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	_, expectedCreateUserError := auth.CreateUser(&repo, username, password, auth.Viewer)
	if expectedCreateUserError == nil {
		t.Fatal("Duplicate user was created")
	}

}

func TestGetCreatedUserByUsername(t *testing.T) {
	// Arrange
	testUser := CreateTestUser()
	repo := infra.NewMemoryAuthRepository()
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := auth.GetUserByUsername(&repo, testUser.Username)
	if getUserError != nil {
		t.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestErrorIsReturnedWhenNoUserByUsername(t *testing.T) {
	// Arrange
	testUser := CreateTestUser()
	repo := infra.NewMemoryAuthRepository()
	// Act
	_, getUserError := auth.GetUserByUsername(&repo, testUser.Username)
	if getUserError == nil {
		t.Fatal("User found when there should not have been")
	}
}

func TestGetCreatedUserById(t *testing.T) {
	// Arrange
	testUser := CreateTestUser()
	repo := infra.NewMemoryAuthRepository()
	_, createUserError := repo.CreateUser(testUser)
	if createUserError != nil {
		t.Fatal("Failed to create initial user")
	}
	// Act
	user, getUserError := auth.GetUserById(&repo, testUser.UserId)
	if getUserError != nil {
		t.Fatal("Failed to get user with matching ID")
	}
	// Assert
	assert.Equal(t, user.UserId, testUser.UserId)
}

func TestErrorIsReturnedWhenNoUserByUserId(t *testing.T) {
	// Arrange
	testUser := CreateTestUser()
	repo := infra.NewMemoryAuthRepository()
	// Act
	_, getUserError := auth.GetUserById(&repo, testUser.UserId)
	if getUserError == nil {
		t.Fatal("User found when there should not have been")
	}
}

func CreateTestUser() auth.User {
	const username = "username"
	const password = "password"
	return auth.NewUser(username, password, auth.Viewer)
}

func TestRevokeUserToken(t *testing.T) {
	// Arrange
	testAuthrepo := infra.NewMemoryAuthRepository()
	const username = "username"
	testAuthrepo.CreateUser(CreateTestUser())
	// Act
	revokeResult := auth.RevokeUserToken(&testAuthrepo, &testAuthrepo, username)
	// Assert
	assert.Equal(t, revokeResult, nil)
}

func TestRevokeUserTokenErrorsWhenUserIsntFound(t *testing.T) {
	// Arrange
	testAuthrepo := infra.NewMemoryAuthRepository()
	const username = "username"
	// Act
	revokeResult := auth.RevokeUserToken(&testAuthrepo, &testAuthrepo, username)
	// Assert
	assert.Equal(t, revokeResult, auth.UserNotFoundError)
}

func TestGrantUserWithCreateClaim(t *testing.T) {
	// Arrange
	const username = "username"
	const password = "password"
	var testUser auth.User = auth.NewUser(username, password, auth.Viewer)
	claim_map := make(map[auth.Claim]auth.Claim)
	claim_map[auth.CREATE_BLOG] = auth.CREATE_BLOG
	repo := infra.NewMemoryAuthRepository()
	repo.CreateUser(testUser)
	// Act
	result := auth.UpdateUserClaims(&repo, testUser.Username, claim_map)
	assert.Equal(t, result, auth.StatusOK)
}
