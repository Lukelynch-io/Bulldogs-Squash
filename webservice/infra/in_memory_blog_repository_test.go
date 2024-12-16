package infra

import (
	"testing"
	"webservice/app/blog"

	"github.com/stretchr/testify/assert"
)

func TestAddBlogPostToBlogRepo(t *testing.T) {
	repo := new(MemoryBlogPostRepository)
	newPost := blog.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	repo.PostBlog(newPost)
}

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	repo := MemoryBlogPostRepository{
		posts: []blog.Post{},
	}
	newPost := blog.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	// Act
	repo.PostBlog(newPost)
	actual := repo.GetBlogs()
	// Assert
	assert.Equal(t, newPost, actual[0])
}
