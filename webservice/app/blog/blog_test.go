package blog_test

import (
	"testing"
	"webservice/app/blog"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestAddBlogPostToBlogRepo(t *testing.T) {
	repo := new(infra.MemoryBlogPostRepository)
	newPost := blog.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	blog.PostBlog(repo, newPost)
}

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryBlogPostRepository)
	newPost := blog.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	// Act
	blog.PostBlog(repo, newPost)
	actual := blog.GetPosts(repo)
	// Assert
	assert.Equal(t, newPost, actual[0])
}
