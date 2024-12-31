package blogpost

import (
	"net/http"
	"webservice/app/auth/claim"
	"webservice/app/blog"
	"webservice/env"
	"webservice/routes/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/blogposts", getBlogPosts)

	router.POST("/blogposts", middleware.BearerTokenMiddleware, addBlogPost)
}

func getBlogPosts(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(blog.IBlogRepository)
	c.IndentedJSON(http.StatusOK, blogRepo.GetBlogs())
}

func addBlogPost(c *gin.Context) {
	blogRepo := c.MustGet(env.BlogRepo).(blog.IBlogRepository)
	userClaims := c.MustGet(env.TokenClaims).([]claim.Claim)
	var newBlogPost blog.Post

	if err := c.BindJSON(&newBlogPost); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	blog.PostBlog(blogRepo, newBlogPost, claim.IntoMap(userClaims))
	c.Status(http.StatusCreated)
}
