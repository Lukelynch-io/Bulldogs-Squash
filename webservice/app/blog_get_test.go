package app_test

import (
	"testing"
	"webservice/app"
	"webservice/domain"
	"webservice/domain/blog_claims"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	var newPost domain.Post = domain.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	authorisedUser := domain.NewUser("username", "password", domain.Viewer)
	repo := new(infra.MemoryBlogPostRepository)
	authorisedUser.Claims[blog_claims.CREATE_BLOG] = blog_claims.CREATE_BLOG
	// Act
	app.PostBlog(repo, newPost, authorisedUser.Claims)
	actual := app.GetPosts(repo)
	// Assert
	assert.Equal(t, newPost, actual[0])
}
