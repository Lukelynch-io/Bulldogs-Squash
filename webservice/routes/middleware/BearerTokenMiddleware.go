package middleware

import (
	"net/http"
	"webservice/app/auth"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

func BearerTokenMiddleware(c *gin.Context) {
	secretKey := c.MustGet(env.SecretKey).([]byte)
	authRepo := c.MustGet(env.AuthRepo).(auth.IAuthRepo)
	authHeader := c.Request.Header.Get("Authorization")
	tokenString, isSuccess := auth.ExtractBearerToken(authHeader)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	_, claims, isSuccess := auth.ValidateToken(tokenString, secretKey)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	parsedUsername := claims.Username
	if parsedUsername != "" {
		_, err := authRepo.GetUserByUsername(parsedUsername)
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		c.Set(env.UsernameToken, parsedUsername)
	} else {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	claimArray := claims.Claims
	c.Set(env.TokenClaims, claimArray)
}
