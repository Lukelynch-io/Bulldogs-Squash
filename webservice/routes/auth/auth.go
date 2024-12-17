package auth

import (
	"net/http"
	"webservice/app/auth"
	blog_claims "webservice/app/blog/claims"

	"github.com/gin-gonic/gin"
)

var SecretKey []byte

func Routes(route *gin.Engine, secretKey []byte) {
	SecretKey = secretKey
	auth := route.Group("/auth")
	{
		auth.POST("/token", requestToken)
	}
}

func requestToken(c *gin.Context) {
	// TODO: Add call to get user claims
	claims := []auth.Claim{
		blog_claims.CREATE_BLOG,
	}
	tokenString, err := auth.GenerateNewToken(SecretKey, claims, "test_claim")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	returnObj := struct {
		Token string
	}{
		Token: tokenString,
	}
	c.IndentedJSON(http.StatusOK, returnObj)
}
