package auth

import (
	"net/http"
	"webservice/app/auth"

	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secret-key")

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestToken)
	}
}

func requestToken(c *gin.Context) {
	tokenString, err := auth.GenerateNewToken(secretKey, "test_claim")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, tokenString)
}
