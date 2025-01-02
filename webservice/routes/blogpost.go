package routes

import (
	"net/http"
	"webservice/app"
	"webservice/domain"
	"webservice/env"

	"github.com/gin-gonic/gin"
)

func LoadBlogPostRoutes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)
	router.POST("/blogposts", BearerTokenMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(domain.IBlogRepository)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(domain.IBlogRepository)
	userClaims := c.MustGet(env.TokenClaims).(domain.ClaimArray)
	var newBlogPost domain.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	app.PostBlog(blogRepo, newBlogPost, userClaims.IntoMap())
	c.Status(http.StatusCreated)
}
