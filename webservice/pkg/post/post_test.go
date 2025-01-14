package post_test

import (
	"mime/multipart"
	"testing"
	"webservice/pkg/auth"
	"webservice/pkg/post"

	"github.com/go-playground/assert/v2"
)

type TestFileStorage struct {
}

func (fs *TestFileStorage) StoreFile(file *multipart.FileHeader) (string, error) {
	return "", nil
}

func TestGetBlogPostFromBlogRepo(t *testing.T) {
	// Arrange
	var newPost post.NewPost = post.NewPost{
		PostTitle:       "Test Title",
		PostDescription: "Test Description",
		FileData:        new(multipart.FileHeader),
	}
	authorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(post.InMemoryPostStorage)
	testFileStorage := new(TestFileStorage)
	authorisedUser.Claims[auth.CREATE_BLOG] = auth.CREATE_BLOG
	// Act
	post.InsertPost(repo, testFileStorage, newPost, authorisedUser.Claims)
	actual := post.GetPosts(repo)
	// Assert
	assert.Equal(t, newPost.PostTitle, actual[0].Title)
}

func TestAddBlogPostToBlogRepo(t *testing.T) {
	// Arrange
	var newPost post.NewPost = post.NewPost{
		PostTitle:       "Test Title",
		PostDescription: "Test Description",
		FileData:        new(multipart.FileHeader),
	}
	authorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(post.InMemoryPostStorage)
	authorisedUser.Claims[auth.CREATE_BLOG] = auth.CREATE_BLOG
	testFileStorage := new(TestFileStorage)

	post.InsertPost(repo, testFileStorage, newPost, authorisedUser.Claims)
}

func TestPostBlogPostAsUnauthorisedUserDoesntWork(t *testing.T) {
	// Arrange
	var newPost post.NewPost = post.NewPost{
		PostTitle:       "Test Title",
		PostDescription: "Test Description",
		FileData:        new(multipart.FileHeader),
	}
	unauthorisedUser := auth.NewUser("username", "password", auth.Viewer)
	repo := new(post.InMemoryPostStorage)
	testFileStorage := new(TestFileStorage)
	// Act
	isPostBlogSuccess, _ := post.InsertPost(repo, testFileStorage, newPost, unauthorisedUser.Claims)
	// Assert
	assert.Equal(t, isPostBlogSuccess, false)
}
