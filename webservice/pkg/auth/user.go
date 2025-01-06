package auth

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type UserId string

type UserRole string

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

type UserDataStorage interface {
	GetUserByUserId(UserId) (*User, error)
	GetUserByUsername(string) (*User, error)
	CreateUser(User) (*User, error)
	GrantUserClaim(UserId, Claim) (bool, error)
}

type UserActionResponses int

const (
	BadRequest        UserActionResponses = http.StatusBadRequest
	UserNotFoundError UserActionResponses = http.StatusNotFound
	UserNotAuthorized UserActionResponses = http.StatusUnauthorized
	StatusNotModified UserActionResponses = http.StatusNotModified
	StatusOK          UserActionResponses = http.StatusOK
)

func CreateUser(repo UserDataStorage, username string, passwordHash string, role UserRole) (*User, error) {
	newUser := NewUser(username, passwordHash, role)
	_, createUserError := repo.CreateUser(newUser)
	if createUserError != nil {
		return nil, createUserError
	}
	return &newUser, nil
}

func GetUserByUsername(repo UserDataStorage, username string) (*User, error) {
	return repo.GetUserByUsername(username)
}

func GetUserById(repo UserDataStorage, userId UserId) (*User, error) {
	return repo.GetUserByUserId(userId)
}

func UpdateUserClaims(authRepo UserDataStorage, targetUsername string, newClaims map[Claim]Claim) UserActionResponses {
	foundUser, err := GetUserByUsername(authRepo, targetUsername)
	if err != nil {
		return StatusNotModified
	}

	for _, claim := range newClaims {
		// FIXME: This is not valid LOGIC
		grantUserClaim(authRepo, foundUser.UserId, foundUser.UserId, claim)
	}
	return StatusOK
}

func grantUserClaim(repo UserDataStorage, actingUserId UserId, targetUserId UserId, newClaim Claim) (bool, error) {
	// TODO: add executing user verification
	actingUser, error := repo.GetUserByUserId(actingUserId)
	if error != nil {
		return false, errors.New("FORBIDDEN")
	}
	if canUserGrantClaim(*actingUser, newClaim) {
		return repo.GrantUserClaim(targetUserId, newClaim)
	}
	return false, errors.New("FORBIDDEN")
}

func canUserGrantClaim(user User, claim Claim) bool {
	return true
}
