package auth

import (
	"net/http"
	"webservice/app/auth"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestToken)
		auth.POST("/user/create", createUser)
	}
}

type RequestTokenObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func requestToken(c *gin.Context) {
	authRepo := c.MustGet("authRepo").(auth.IAuthRepo)
	// TODO: Add call to get user claims
	var userDetails RequestTokenObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	secretKey := c.MustGet("secretKey").([]byte)
	tokenString, err := auth.HandleUserAuth(authRepo, secretKey, userDetails.Username, userDetails.Password)
	if err != 0 {
		c.Status(err)
		return
	}

	returnObj := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}
	c.IndentedJSON(http.StatusOK, returnObj)
}
