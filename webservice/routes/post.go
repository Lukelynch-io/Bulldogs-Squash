package routes

import (
	"net/http"
	"webservice/env"
	"webservice/pkg/auth"
	"webservice/pkg/filestore"
	"webservice/pkg/post"

	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(router *gin.Engine) {
	router.GET("/blogposts", getPosts)
	router.POST("/blogposts", BearerTokenMiddleware, addPost)
}

func getPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	trimRequest := c.Query("trimmed")
	trimRequestBool := false
	if trimRequest == "true" {
		trimRequestBool = true
	}
	posts, err := post.GetPosts(blogRepo, trimRequestBool)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func addPost(c *gin.Context) {
	postRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	userClaims := c.MustGet(env.TokenClaims).(auth.ClaimArray)
	fileStorage := c.MustGet(env.FileStorage).(filestore.Filestore)
	var newBlogPost post.NewPost

	if err := c.Bind(&newBlogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.InsertPost(postRepo, fileStorage, newBlogPost, userClaims.IntoMap())
	c.Status(http.StatusCreated)
}
