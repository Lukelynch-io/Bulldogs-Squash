package blogpost

import (
	"fmt"
	"net/http"
	"webservice/app/blog"

	"github.com/gin-gonic/gin"
)

func CheckAuthorizationMiddleware(c *gin.Context) {
	fmt.Printf("Header: %s", c.Request.Header.Get("Authorization"))

}

func Routes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)

	router.POST("/blogposts", CheckAuthorizationMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet("blogRepo").(blog.IBlogRepository)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	// blogRepo := c.MustGet("blogRepo").(blog.IBlogRepository)
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		return
	}
	// blog.PostBlog(blogRepo, newBlogPost)
	c.Status(http.StatusCreated)
}
