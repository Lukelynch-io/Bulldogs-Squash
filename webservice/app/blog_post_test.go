package app_test

import (
	"testing"
	"webservice/app"
	"webservice/domain"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestAddBlogPostToBlogRepo(t *testing.T) {
	// Arrange
	var newPost domain.Post = domain.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	authorisedUser := domain.NewUser("username", "password", domain.Viewer)
	repo := new(infra.MemoryBlogPostRepository)
	authorisedUser.Claims[domain.CREATE_BLOG] = domain.CREATE_BLOG

	app.PostBlog(repo, newPost, authorisedUser.Claims)
}

func TestPostBlogPostAsUnauthorisedUserDoesntWork(t *testing.T) {
	// Arrange
	var newPost domain.Post = domain.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	unauthorisedUser := domain.NewUser("username", "password", domain.Viewer)
	repo := new(infra.MemoryBlogPostRepository)
	// Act
	isPostBlogSuccess, _ := app.PostBlog(repo, newPost, unauthorisedUser.Claims)
	// Assert
	assert.Equal(t, isPostBlogSuccess, false)
}
