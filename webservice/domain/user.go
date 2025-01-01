package domain

import (
	"github.com/google/uuid"
)

type UserId string

type userRole string

const AdminRole = userRole("admin")
const Poster = userRole("poster")
const Viewer = userRole("viewer")

type User struct {
	UserId       UserId
	Username     string
	PasswordHash string
	Claims       map[Claim]Claim
	Role         userRole
}

func NewUser(username string, passwordHash string) User {
	return User{
		UserId:       UserId(uuid.New().String()),
		Username:     username,
		PasswordHash: passwordHash,
		Claims:       make(map[Claim]Claim),
	}
}
