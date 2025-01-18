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
	router.GET("/post/:postId", getPost)
	router.GET("/posts", getPosts)
	router.POST("/posts", BearerTokenMiddleware, addPost)
}

func getPost(c *gin.Context) {
	postRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	postId := c.Param("postId")
	if postId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	post, err := post.GetPost(postRepo, postId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, post)
}

func getPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	trimResponse := c.Query("trimmed") == "true"
	posts, err := post.GetPosts(blogRepo, trimResponse)
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
