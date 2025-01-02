package routes

import (
	"net/http"
	"webservice/app"
	"webservice/domain"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

type createUserRequestObj struct {
	Username     string          `json:"username"`
	PasswordHash string          `json:"passwordHash"`
	Role         domain.UserRole `json:"role"`
}

type createUserErrorResponse struct {
	Error string `json:"error"`
}

func createUser(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(domain.IAuthRepo)
	var userDetails createUserRequestObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, createUserError := app.CreateUser(authRepo, userDetails.Username, userDetails.PasswordHash, userDetails.Role)
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
	Username string                        `json:"username"`
	Claims   map[domain.Claim]domain.Claim `json:"claims"`
}

func updateUserClaims(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(domain.IAuthRepo)
	var newUserClaims updateUserClaimsRequestObj

	if err := c.BindJSON(&newUserClaims); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	app.UpdateUserClaims(authRepo, newUserClaims.Username, newUserClaims.Claims)
	c.Status(http.StatusOK)

}
