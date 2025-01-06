package infra

import (
	"strconv"
	"webservice/pkg/post"
)

type MemoryBlogStorage struct {
	posts []post.Post
}

func (repo *MemoryBlogStorage) getNewPostId() (string, error) {
	var latestPost = repo.posts[len(repo.posts)-1]
	i, err := strconv.Atoi(latestPost.ID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(i + 1), nil
}

func (repo *MemoryBlogStorage) GetBlogs() []post.Post {
	return repo.posts
}

func (repo *MemoryBlogStorage) PostBlog(post post.Post) (bool, error) {
	var newId, err = repo.getNewPostId()
	if err != nil {
		return false, err
	}
	post.ID = newId
	repo.posts = append(repo.posts, post)
	return true, nil
}

func (repo *MemoryBlogStorage) LoadAllPosts(posts []post.Post) {
	repo.posts = posts
}
