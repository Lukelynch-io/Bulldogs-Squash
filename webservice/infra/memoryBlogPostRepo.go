package infra

import (
	"strconv"
	"webservice/domain"
)

type MemoryBlogPostRepository struct {
	posts []domain.Post
}

func (repo *MemoryBlogPostRepository) getNewPostId() (string, error) {
	var latestPost = repo.posts[len(repo.posts)-1]
	i, err := strconv.Atoi(latestPost.ID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(i + 1), nil
}

func (repo *MemoryBlogPostRepository) GetBlogs() []domain.Post {
	return repo.posts
}

func (repo *MemoryBlogPostRepository) PostBlog(post domain.Post) (bool, error) {
	if len(repo.posts) == 0 {
		repo.posts = []domain.Post{post}
		return true, nil
	}
	var newId, err = repo.getNewPostId()
	if err != nil {
		return false, err
	}
	post.ID = newId
	repo.posts = append(repo.posts, post)
	return true, nil
}

func (repo *MemoryBlogPostRepository) LoadAllPosts(posts []domain.Post) {
	repo.posts = posts
}
