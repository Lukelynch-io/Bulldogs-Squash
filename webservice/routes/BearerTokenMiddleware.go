package routes

import (
	"net/http"
	"webservice/env"
	"webservice/pkg/auth"

	"github.com/gin-gonic/gin"
)

func BearerTokenMiddleware(c *gin.Context) {
	secretKey := c.MustGet(env.SecretKey).([]byte)
	authRepo := c.MustGet(env.AuthRepo).(auth.UserDataStorage)
	authHeader := c.Request.Header.Get("Authorization")
	tokenString, isSuccess := auth.ExtractBearerToken(authHeader)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	claims, validationError := auth.ValidateToken(tokenString, secretKey)
	if validationError != nil {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	parsedUserId := claims.Subject
	if parsedUserId != "" {
		user, err := auth.GetUserById(authRepo, auth.UserId(parsedUserId))
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		c.Set(env.User, *user)
	} else {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	claimArray := claims.Claims
	c.Set(env.TokenClaims, claimArray)
}
