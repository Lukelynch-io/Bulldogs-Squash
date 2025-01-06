package routes

import (
	"net/http"
	"webservice/env"
	"webservice/pkg/auth"
	"webservice/pkg/post"

	"github.com/gin-gonic/gin"
)

func LoadBlogPostRoutes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", BearerTokenMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	userClaims := c.MustGet(env.TokenClaims).(auth.ClaimArray)
	var newBlogPost post.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	post.InsertPost(blogRepo, newBlogPost, userClaims.IntoMap())
	c.Status(http.StatusCreated)
}
