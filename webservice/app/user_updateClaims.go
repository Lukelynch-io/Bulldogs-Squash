package app

import (
	"errors"
	"webservice/domain"
)

func UpdateUserClaims(authRepo domain.IAuthRepo, targetUsername string, newClaims map[domain.Claim]domain.Claim) {
	foundUser, err := GetUserByUsername(authRepo, targetUsername)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	for _, claim := range newClaims {
		// FIXME: This is not valid LOGIC
		grantUserClaim(authRepo, foundUser.UserId, foundUser.UserId, claim)
	}
}

func grantUserClaim(repo domain.IAuthRepo, actingUserId domain.UserId, targetUserId domain.UserId, newClaim domain.Claim) (bool, error) {
	// TODO: add executing user verification
	actingUser, error := repo.GetUserByUserId(actingUserId)
	if error != nil {
		return false, errors.New("FORBIDDEN")
	}
	if canUserGrantClaim(actingUser, newClaim) {
		return repo.GrantUserClaim(targetUserId, newClaim)
	}
	return false, errors.New("FORBIDDEN")
}

func canUserGrantClaim(user domain.User, claim domain.Claim) bool {
	return true
}
