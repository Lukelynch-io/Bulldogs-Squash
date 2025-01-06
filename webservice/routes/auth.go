package routes

import (
	"net/http"
	"webservice/env"
	"webservice/pkg/auth"

	"github.com/gin-gonic/gin"
)

func LoadAuthRoutes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestUserToken)
		auth.DELETE("/token", revokeUserToken)
	}

}

type RequestTokenObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func requestUserToken(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(auth.UserDataStorage)
	secretKey := c.MustGet(env.SecretKey).([]byte)
	// TODO: Add call to get user claims
	var userDetails RequestTokenObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	tokenString, resultType := auth.AuthenticateUser(authRepo, secretKey, userDetails.Username, userDetails.Password)
	if int(resultType) != http.StatusOK {
		c.Status(int(resultType))
		return
	}

	returnObj := struct {
		Token auth.TokenString `json:"token"`
	}{
		Token: *tokenString,
	}
	c.IndentedJSON(http.StatusOK, returnObj)
}

func revokeUserToken(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(auth.UserDataStorage)
	tokenRepo := c.MustGet(env.TokenStorage).(auth.TokenStorage)
	username := c.Query("username")

	err := auth.RevokeUserToken(authRepo, tokenRepo, username)
	if err != nil {
		if int(*err) == int(auth.UserNotFoundError) {
			c.Status(http.StatusNotFound)
			return
		}
	}
	c.Status(http.StatusOK)
}
