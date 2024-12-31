package blog_test

import (
	"testing"
	"webservice/app/auth"
	"webservice/app/blog"
	"webservice/app/blog/blog_claims"
	"webservice/infra"

	"github.com/go-playground/assert/v2"
)

var newPost blog.Post = blog.Post{
	ID:          "1",
	Title:       "Test Title",
	Description: "Test Description",
	ImageUrl:    "imageUrl",
}
var authorisedUser auth.User = auth.NewUser("username", "password")
var unauthorisedUser auth.User = auth.NewUser("username", "password")

func TestAddBlogPostToBlogRepo(t *testing.T) {
	repo := new(infra.MemoryBlogPostRepository)
	authorisedUser.Claims[blog_claims.CREATE_BLOG] = blog_claims.CREATE_BLOG

	blog.PostBlog(repo, newPost, authorisedUser.Claims)
}

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryBlogPostRepository)
	authorisedUser.Claims[blog_claims.CREATE_BLOG] = blog_claims.CREATE_BLOG
	// Act
	blog.PostBlog(repo, newPost, authorisedUser.Claims)
	actual := blog.GetPosts(repo)
	// Assert
	assert.Equal(t, newPost, actual[0])
}

func TestPostBlogPostAsUnauthorisedUserDoesntWork(t *testing.T) {
	// Arrange
	repo := new(infra.MemoryBlogPostRepository)
	// Act
	isPostBlogSuccess, _ := blog.PostBlog(repo, newPost, unauthorisedUser.Claims)
	// Assert
	assert.Equal(t, isPostBlogSuccess, false)
}
