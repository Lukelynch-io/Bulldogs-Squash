package infra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBlogPostToBlogRepo(t *testing.T) {
	repo := new(BlogPostRepository)
	repo.PostBlog("This is a test blog post")
}

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	repo := BlogPostRepository{
		posts: []string{},
	}
	test_string := "This is a test blog post"
	// Act
	repo.PostBlog(test_string)
	actual := repo.GetBlogs()
	// Assert
	assert.Equal(t, test_string, actual[0])
}
