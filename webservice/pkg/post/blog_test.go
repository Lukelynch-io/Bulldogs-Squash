package post_test

import (
	"testing"
	"webservice/infra"
	"webservice/pkg/auth"
	"webservice/pkg/post"

	"github.com/go-playground/assert/v2"
)

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	var newPost post.Post = post.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	authorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(infra.MemoryBlogStorage)
	authorisedUser.Claims[auth.CREATE_BLOG] = auth.CREATE_BLOG
	// Act
	post.PostBlog(repo, newPost, authorisedUser.Claims)
	actual := post.GetPosts(repo)
	// Assert
	assert.Equal(t, newPost, actual[0])
}

func TestAddBlogPostToBlogRepo(t *testing.T) {
	// Arrange
	var newPost post.Post = post.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	authorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(infra.MemoryBlogStorage)
	authorisedUser.Claims[auth.CREATE_BLOG] = auth.CREATE_BLOG

	post.PostBlog(repo, newPost, authorisedUser.Claims)
}

func TestPostBlogPostAsUnauthorisedUserDoesntWork(t *testing.T) {
	// Arrange
	var newPost post.Post = post.Post{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	unauthorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(infra.MemoryBlogStorage)
	// Act
	isPostBlogSuccess, _ := post.PostBlog(repo, newPost, unauthorisedUser.Claims)
	// Assert
	assert.Equal(t, isPostBlogSuccess, false)
}
