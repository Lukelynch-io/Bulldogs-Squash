package auth

import (
	"net/http"
	"webservice/app"
	"webservice/app/auth"
	"webservice/domain"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestToken)
		userAuth := auth.Group("/user")
		{
			userAuth.POST("/create", createUser)
			userAuth.PATCH("/claims", updateUserClaims)
		}
	}

}

type RequestTokenObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func requestToken(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(domain.IAuthRepo)
	secretKey := c.MustGet(env.SecretKey).([]byte)
	// TODO: Add call to get user claims
	var userDetails RequestTokenObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	tokenString, err := app.AuthenticateUser(authRepo, secretKey, userDetails.Username, userDetails.Password)
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
