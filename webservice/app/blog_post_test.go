package app_test

import (
	"testing"
	"webservice/app"
	"webservice/app/blog/blog_claims"
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
	var authorisedUser domain.User = domain.NewUser("username", "password")
	repo := new(infra.MemoryBlogPostRepository)
	authorisedUser.Claims[blog_claims.CREATE_BLOG] = blog_claims.CREATE_BLOG

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
	var unauthorisedUser domain.User = domain.NewUser("username", "password")
	repo := new(infra.MemoryBlogPostRepository)
	// Act
	isPostBlogSuccess, _ := app.PostBlog(repo, newPost, unauthorisedUser.Claims)
	// Assert
	assert.Equal(t, isPostBlogSuccess, false)
}
