package blog

import (
	"errors"
	"webservice/app/auth"
	"webservice/app/blog/blog_claims"
)

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type IBlogRepository interface {
	GetBlogs() []Post
	PostBlog(Post) (bool, error)
	LoadAllPosts([]Post)
}

func GetPosts(repo IBlogRepository) []Post {
	return repo.GetBlogs()
}

func PostBlog(repo IBlogRepository, post Post, user auth.User) (bool, error) {
	if user.Claims[blog_claims.CREATE_BLOG] == blog_claims.CREATE_BLOG {
		repo.PostBlog(post)
		return true, nil
	}
	return false, errors.New("User does not have permission to perform this action")
}
