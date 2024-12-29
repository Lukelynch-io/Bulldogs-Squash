package auth

import (
	"net/http"
	"webservice/app/auth"

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
	authRepo := c.MustGet("authRepo").(auth.IAuthRepo)
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