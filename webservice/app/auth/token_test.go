package auth_test

import (
	"testing"
	"webservice/app/auth"
	"webservice/app/auth/claim"

	"github.com/golang-jwt/jwt"
)

func TestGenerateAndValidateToken(t *testing.T) {
	var secretKey []byte
	const username string = "username"
	var claims []claim.Claim = []claim.Claim{}

	tokenString, err := auth.GenerateNewToken(secretKey, claims, username)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	token, isSuccess := auth.ValidateToken(tokenString, secretKey)
	if !isSuccess {
		t.Fail()
	}
	if parsedClaims, ok := token.Claims.(jwt.MapClaims); !ok {
		t.Log(err)
		t.Fail()
	} else {
		t.Logf("Raw Token: %s", parsedClaims["username"])

	}
}
