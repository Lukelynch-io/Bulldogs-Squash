package blogpost

import (
	"net/http"
	"webservice/app/blog"

	"github.com/gin-gonic/gin"
)

var iBlogRepo blog.IBlogRepository

func Routes(router *gin.Engine, repo blog.IBlogRepository) {
	iBlogRepo = repo
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, iBlogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		return
	}
	iBlogRepo.PostBlog(newBlogPost)
	c.Status(http.StatusCreated)
}
