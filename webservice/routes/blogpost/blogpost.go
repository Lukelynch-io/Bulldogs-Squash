package blogpost

import (
	"net/http"
	"webservice/app/blog"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet("blogRepo").(blog.IBlogRepository)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	blogRepo := c.MustGet("blogRepo").(blog.IBlogRepository)
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		return
	}
	blogRepo.PostBlog(newBlogPost)
	c.Status(http.StatusCreated)
}
