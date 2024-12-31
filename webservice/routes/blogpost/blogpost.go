package blogpost

import (
	"net/http"
	"webservice/app/auth"
	"webservice/app/blog"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

func CheckAuthorizationMiddleware(c *gin.Context) {
	secretKey := c.MustGet(env.SecretKey).([]byte)
	authHeader := c.Request.Header.Get("Authorization")
	//TODO: Figure out how this should be used
	tokenString, isSuccess := auth.ExtractBearerToken(authHeader)
	if !isSuccess {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	auth.ValidateToken(tokenString, secretKey)

}

func Routes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)

	router.POST("/blogposts", CheckAuthorizationMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(blog.IBlogRepository)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	// blogRepo := c.MustGet(env.BlogRepo).(blog.IBlogRepository)
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		return
	}
	// blog.PostBlog(blogRepo, newBlogPost)
	c.Status(http.StatusCreated)
}
