package blogpost

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
