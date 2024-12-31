package blogpost

import (
	"fmt"
	"net/http"
	"webservice/app/auth"
	"webservice/app/auth/claim"
	"webservice/app/blog"
	"webservice/env"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func BearerTokenMiddleware(c *gin.Context) {
	secretKey := c.MustGet(env.SecretKey).([]byte)
	authRepo := c.MustGet(env.AuthRepo).(auth.IAuthRepo)
	authHeader := c.Request.Header.Get("Authorization")
	// TODO: Figure out how this should be used
	tokenString, isSuccess := auth.ExtractBearerToken(authHeader)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	token, isSuccess := auth.ValidateToken(tokenString, secretKey)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	parsedClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
	// TODO: Need to ensure that the tokens
	parsedUsername := parsedClaims["username"].(string)
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
	var claimArray []claim.Claim
	extractedInterfaceArray, ok := parsedClaims["claims"].([]interface{})
	for _, obj := range extractedInterfaceArray {
		claimArray = append(claimArray, claim.Claim(obj.(string)))
	}
	if !ok {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	fmt.Println("We got here...", claimArray)
	c.Set(env.TokenClaims, claimArray)
}

func Routes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)

	router.POST("/blogposts", BearerTokenMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(blog.IBlogRepository)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(blog.IBlogRepository)
	userClaims := c.MustGet(env.TokenClaims).([]claim.Claim)
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	blog.PostBlog(blogRepo, newBlogPost, claim.IntoMap(userClaims))
	c.Status(http.StatusCreated)
}
