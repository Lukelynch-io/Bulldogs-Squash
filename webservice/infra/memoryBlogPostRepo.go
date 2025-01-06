package infra

import (
	"webservice/pkg/post"

	"github.com/google/uuid"
)

type MemoryBlogStorage struct {
	posts []post.Post
}

func (repo *MemoryBlogStorage) GetBlogs() []post.Post {
	return repo.posts
}

func (repo *MemoryBlogStorage) PostBlog(post post.Post) (bool, error) {
	post.ID = uuid.NewString()
	repo.posts = append(repo.posts, post)
	return true, nil
}

func (repo *MemoryBlogStorage) LoadAllPosts(posts []post.Post) {
	repo.posts = posts
}
