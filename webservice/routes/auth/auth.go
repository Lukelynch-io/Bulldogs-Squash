package auth

import (
	"net/http"
	"webservice/app/auth"

	"github.com/gin-gonic/gin"
)

var SecretKey []byte
var AuthRepo auth.IAuthRepo

func Routes(route *gin.Engine, secretKey []byte, authRepo auth.IAuthRepo) {
	SecretKey = secretKey
	AuthRepo = authRepo
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestToken)
	}
}

type requestObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func requestToken(c *gin.Context) {
	// TODO: Add call to get user claims
	var userDetails requestObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	tokenString, err := auth.HandleUserAuth(SecretKey, userDetails.Username, userDetails.Password)
	if err != 0 {
		c.Status(err)
	}

	returnObj := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}
	c.IndentedJSON(http.StatusOK, returnObj)
}
