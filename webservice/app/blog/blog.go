package blog

import (
	"errors"
	"webservice/app/auth/claim"
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

func PostBlog(blogRepo IBlogRepository, post Post, userClaims map[claim.Claim]claim.Claim) (bool, error) {
	if userClaims[blog_claims.CREATE_BLOG] == blog_claims.CREATE_BLOG {
		blogRepo.PostBlog(post)
		return true, nil
	}
	return false, errors.New("User does not have permission to perform this action")
}
