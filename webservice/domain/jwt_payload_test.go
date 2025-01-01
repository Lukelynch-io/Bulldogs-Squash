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
	tokenString, err := domain.GenerateNewToken(secretKey, claims, username, &testTime)
	t.Log(tokenString)
	if err != nil {
		t.Fail()
		return
	}
	claims, err := domain.ValidateToken(tokenString, secretKey)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(claims)
	assert.Equal(t, claims.Username, username)
}
