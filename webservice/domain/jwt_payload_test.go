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

func TestGenerateAndValidateToken(t *testing.T) {
	const username = "username"
	const password = "password"
	user := domain.NewUser(username, password, domain.Viewer)
	tokenString, err := domain.GenerateNewToken(secretKey, claims, user, &testTime)
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
	assert.Equal(t, claims.Subject, string(user.UserId))
}
