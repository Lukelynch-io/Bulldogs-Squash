package auth

import (
	"webservice/app/auth/claim"

	"github.com/google/uuid"
)

type UserId string

type User struct {
	UserId       UserId
	Username     string
	PasswordHash string
	Claims       map[claim.Claim]claim.Claim
}

func NewUser(username string, passwordHash string) User {
	return User{
		UserId:       UserId(uuid.New().String()),
		Username:     username,
		PasswordHash: passwordHash,
		Claims:       make(map[claim.Claim]claim.Claim),
	}
}
