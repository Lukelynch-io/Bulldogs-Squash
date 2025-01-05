package routes

import (
	"net/http"
	"webservice/app"
	"webservice/domain"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

func BearerTokenMiddleware(c *gin.Context) {
	secretKey := c.MustGet(env.SecretKey).([]byte)
	authRepo := c.MustGet(env.AuthRepo).(domain.IAuthRepo)
	authHeader := c.Request.Header.Get("Authorization")
	tokenString, isSuccess := domain.ExtractBearerToken(authHeader)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	claims, validationError := domain.ValidateToken(tokenString, secretKey)
	if validationError != nil {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	parsedUserId := claims.Subject
	if parsedUserId != "" {
		user, err := app.GetUserById(authRepo, domain.UserId(parsedUserId))
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
