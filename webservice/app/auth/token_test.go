package auth_test

import (
	"testing"
	"time"
	"webservice/app/auth"
	"webservice/app/auth/claim"

	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt"
)

var secretKey []byte
var claims []claim.Claim = []claim.Claim{}

var testTime = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

const expectedToken = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbXMiOltdLCJleHAiOjE3MzU2ODk2MDAsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.X_HoDH0jF7e2UHWAN8iaE3LRd7x2suIdvqOnNp0rshY`

func TestGenerateAndValidateToken(t *testing.T) {
	tokenString, err := auth.GenerateNewToken(secretKey, claims, username, &testTime)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, tokenString, expectedToken)
}

func TestValidateToken(t *testing.T) {
	token, isSuccess := auth.ValidateToken(expectedToken, secretKey)
	if !isSuccess {
		t.Fail()
		return
	}
	if parsedClaims, ok := token.Claims.(jwt.MapClaims); !ok {
		t.Fail()
		return
	} else {
		t.Logf("Raw Token: %s", parsedClaims["username"])

	}
}
