package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestToken)
	}
}

func requestToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": "test_username",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, tokenString)
}
