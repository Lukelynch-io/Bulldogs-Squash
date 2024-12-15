package main

import (
	"net/http"
	"webservice/app/blog"
	"webservice/app/infra_interfaces"
	"webservice/infra"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var iblogPostRepository infra_interfaces.IBlogRepository

func getBlogPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, iblogPostRepository.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		return
	}
	iblogPostRepository.PostBlog(newBlogPost)
	c.Status(http.StatusCreated)
}

func main() {
	iblogPostRepository = &infra.MemoryBlogPostRepository{}
	iblogPostRepository.LoadAllPosts(testdata.LoadDummyData())

	router := gin.Default()
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", addBlogPost)

	router.Run("localhost:8080")
}
