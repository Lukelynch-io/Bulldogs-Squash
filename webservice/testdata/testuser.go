package testdata

import (
	"crypto/sha256"
	// "webservice/pkg/auth"
)

func CreateTestUsers() {
	var password1 = "Password1"
	h := sha256.New()
	h.Write([]byte(password1))

	// var _ = auth.NewUser("hello", string(h.Sum(nil)), auth.AdminRole)
}
