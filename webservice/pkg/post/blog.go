package post

import (
	"errors"
	"webservice/pkg/auth"
)

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type PostStorage interface {
	GetBlogs() []Post
	InsertPost(Post) (bool, error)
	LoadAllPosts([]Post)
}

func GetPosts(repo PostStorage) []Post {
	return repo.GetBlogs()
}

func InsertPost(blogRepo PostStorage, post Post, actingUserClaims map[auth.Claim]auth.Claim) (bool, error) {
	if actingUserClaims[auth.CREATE_BLOG] == auth.CREATE_BLOG {
		blogRepo.InsertPost(post)
		return true, nil
	}
	return false, errors.New("User does not have permission to perform this action")
}
