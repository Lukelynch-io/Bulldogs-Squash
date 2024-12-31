package auth_test

import (
	"testing"
	"time"
	"webservice/app/auth"
	"webservice/app/auth/claim"

	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte
var claims []claim.Claim = []claim.Claim{}

var testTime = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

const expectedToken = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lIiwiY2xhaW1zIjpbXSwiZXhwIjoxNzM1Njg5NjAwfQ.hHy-nkmmQkgaYAQs5GEpU1kdTv4vqN2zCFGT0MZcB_M`

func TestGenerateAndValidateToken(t *testing.T) {
	tokenString, err := auth.GenerateNewToken(secretKey, claims, username, &testTime)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, tokenString, expectedToken)
}

func TestValidateToken(t *testing.T) {
	token, claims, isSuccess := auth.ValidateToken(expectedToken, secretKey)
	if !isSuccess {
		t.Fail()
		return
	}
	t.Log(claims)
	if parsedClaims, ok := token.Claims.(jwt.MapClaims); !ok {
		t.Fail()
		return
	} else {
		t.Logf("Raw Token: %s", parsedClaims["username"])

	}
}

func TestGetClaimFromToken(t *testing.T) {
	token, claims, isSuccess := auth.ValidateToken(expectedToken, secretKey)
	if !isSuccess {
		t.Fail()
		return
	}
	t.Log(claims)
	parsedClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fail()
		return
	}
	assert.Equal(t, parsedClaims["username"], username)
}
