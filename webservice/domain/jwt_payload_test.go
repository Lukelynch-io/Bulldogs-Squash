package domain_test

import (
	"testing"
	"time"
	"webservice/domain"

	"github.com/go-playground/assert/v2"
)

var secretKey []byte
var claims []domain.Claim = []domain.Claim{}

var testTime = time.Date(3000, time.January, 1, 0, 0, 0, 0, time.UTC)

func CreateTestUser() domain.User {
	const username = "username"
	const password = "password"
	return domain.NewUser(username, password, domain.Viewer)
}

func TestGenerateAndValidateToken(t *testing.T) {
	testUser := CreateTestUser()
	tokenString, err := domain.GenerateNewToken(secretKey, claims, testUser, &testTime)
	t.Log(tokenString)
	if err != nil {
		t.Fail()
		return
	}
	claims, err := domain.ValidateToken(string(*tokenString), secretKey)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(claims)
	assert.Equal(t, claims.Subject, string(testUser.UserId))
}

func TestValidationErrorWhenTokenIsExpired(t *testing.T) {
	testUser := CreateTestUser()
	expiredTime := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	tokenString, err := domain.GenerateNewToken(secretKey, claims, testUser, &expiredTime)
	t.Log(tokenString)
	if err != nil {
		t.Fail()
		return
	}
	claims, err := domain.ValidateToken(string(*tokenString), secretKey)
	if err != nil && err.Error() == "Token has expired" {
		return
	}
	t.Log(claims)
	t.Fail()

}
