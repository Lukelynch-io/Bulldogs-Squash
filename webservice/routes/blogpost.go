package routes

import (
	"mime/multipart"
	"net/http"
	"webservice/env"
	"webservice/pkg/post"

	"github.com/gin-gonic/gin"
)

func LoadBlogPostRoutes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

type PostWithFile struct {
	PostTitle       string                `form:"postTitle"`
	PostDescription string                `form:"postDescription"`
	FileData        *multipart.FileHeader `form:"imageFile"`
}

func addBlogPost(c *gin.Context) {
	// blogRepo := c.MustGet(env.PostRepo).(post.PostStorage)
	// userClaims := c.MustGet(env.TokenClaims).(auth.ClaimArray)
	var newBlogPost PostWithFile

	if err := c.ShouldBind(&newBlogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.SaveUploadedFile(newBlogPost.FileData, "/Users/lukelynch/Desktop/hello.png")
	c.Status(http.StatusCreated)
}
