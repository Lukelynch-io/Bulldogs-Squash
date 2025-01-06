package post

import (
	"github.com/google/uuid"
)

// An In memory post storage that does not persist after program shutsdown
type InMemoryPostStorage struct {
	posts []Post
}

func (repo *InMemoryPostStorage) GetBlogs() []Post {
	return repo.posts
}

func (repo *InMemoryPostStorage) PostBlog(post Post) (bool, error) {
	post.ID = uuid.NewString()
	repo.posts = append(repo.posts, post)
	return true, nil
}

func (repo *InMemoryPostStorage) LoadAllPosts(posts []Post) {
	repo.posts = posts
}
