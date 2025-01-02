package domain

import (
	"github.com/google/uuid"
)

type UserId string

type UserRole string

const AdminRole = UserRole("admin")
const Poster = UserRole("poster")
const Viewer = UserRole("viewer")

type User struct {
	UserId       UserId
	Username     string
	PasswordHash string
	Claims       ClaimMap
	Role         UserRole
}

func NewUser(username string, passwordHash string, role UserRole) User {
	return User{
		UserId:       UserId(uuid.New().String()),
		Username:     username,
		PasswordHash: passwordHash,
		Claims:       make(ClaimMap),
		Role:         role,
	}
}
