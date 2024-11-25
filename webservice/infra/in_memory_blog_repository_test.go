package infra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBlogPostToBlogRepo(t *testing.T) {
	repo := new(MemoryBlogPostRepository)
	repo.PostBlog("This is a test blog post")
}

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	repo := MemoryBlogPostRepository{
		posts: []string{},
	}
	test_string := "This is a test blog post"
	// Act
	repo.PostBlog(test_string)
	actual := repo.GetBlogs()
	// Assert
	assert.Equal(t, test_string, actual[0])
}
