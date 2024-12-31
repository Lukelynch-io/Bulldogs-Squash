package auth

import (
	"net/http"
	"webservice/app/auth"
	"webservice/app/auth/claim"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

type createUserRequestObj struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

type createUserErrorResponse struct {
	Error string `json:"error"`
}

func createUser(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(auth.IAuthRepo)
	var userDetails createUserRequestObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, createUserError := auth.CreateUser(authRepo, userDetails.Username, userDetails.PasswordHash)
	if createUserError != nil {
		c.IndentedJSON(http.StatusInternalServerError, createUserErrorResponse{
			Error: createUserError.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
	return
}

type updateUserClaimsRequestObj struct {
	Username string                      `json:"username"`
	Claims   map[claim.Claim]claim.Claim `json:"claims"`
}

func updateUserClaims(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(auth.IAuthRepo)
	var newUserClaims updateUserClaimsRequestObj

	if err := c.BindJSON(&newUserClaims); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	foundUser, err := auth.GetUserByUsername(authRepo, newUserClaims.Username)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	for _, claim := range newUserClaims.Claims {
		auth.GrantUserClaim(authRepo, foundUser.UserId, claim)
	}
	c.Status(http.StatusOK)

}
