package routes

import (
	"net/http"
	"webservice/env"
	"webservice/pkg/auth"
	"webservice/pkg/filestore"
	"webservice/pkg/post"

	"github.com/gin-gonic/gin"
)

func LoadBlogPostRoutes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", BearerTokenMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	c.IndentedJSON(http.StatusOK, post.GetPosts(blogRepo))
}

func addBlogPost(c *gin.Context) {
	postRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	userClaims := c.MustGet(env.TokenClaims).(auth.ClaimArray)
	var newBlogPost post.NewPost

	if err := c.Bind(&newBlogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filestore.SetFilestoreLocation("/Users/lukelynch/Desktop/")
	post.InsertPost(postRepo, newBlogPost, userClaims.IntoMap())
	c.Status(http.StatusCreated)
}
