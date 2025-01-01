package auth_test

import (
	"testing"
	"time"
	"webservice/app/auth"
	"webservice/app/auth/claim"

	"github.com/go-playground/assert/v2"
)

var secretKey []byte
var claims []claim.Claim = []claim.Claim{}

var testTime = time.Date(3000, time.January, 1, 0, 0, 0, 0, time.UTC)

func TestGenerateAndValidateToken(t *testing.T) {
	tokenString, err := auth.GenerateNewToken(secretKey, claims, username, &testTime)
	t.Log(tokenString)
	if err != nil {
		t.Fail()
		return
	}
	claims, err := auth.ValidateToken(tokenString, secretKey)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(claims)
	assert.Equal(t, claims.Username, username)
}
