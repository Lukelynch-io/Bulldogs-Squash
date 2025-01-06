package post_test

import (
	"testing"
	"webservice/pkg/auth"
	"webservice/pkg/post"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	var newPost post.Post = post.Post{
		ID:          uuid.NewString(),
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	authorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(post.InMemoryPostStorage)
	authorisedUser.Claims[auth.CREATE_BLOG] = auth.CREATE_BLOG
	// Act
	post.InsertPost(repo, newPost, authorisedUser.Claims)
	actual := post.GetPosts(repo)
	// Assert
	assert.Equal(t, newPost, actual[0])
}

func TestAddBlogPostToBlogRepo(t *testing.T) {
	// Arrange
	var newPost post.Post = post.Post{
		ID:          uuid.NewString(),
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	authorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(post.InMemoryPostStorage)
	authorisedUser.Claims[auth.CREATE_BLOG] = auth.CREATE_BLOG

	post.InsertPost(repo, newPost, authorisedUser.Claims)
}

func TestPostBlogPostAsUnauthorisedUserDoesntWork(t *testing.T) {
	// Arrange
	var newPost post.Post = post.Post{
		ID:          uuid.NewString(),
		Title:       "Test Title",
		Description: "Test Description",
		ImageUrl:    "imageUrl",
	}
	unauthorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(post.InMemoryPostStorage)
	// Act
	isPostBlogSuccess, _ := post.InsertPost(repo, newPost, unauthorisedUser.Claims)
	// Assert
	assert.Equal(t, isPostBlogSuccess, false)
}
