package infra

import (
	"strconv"
	"webservice/app/blog"
)

type MemoryBlogPostRepository struct {
	posts []blog.Post
}

func (repo *MemoryBlogPostRepository) getNewPostId() (string, error) {
	var latestPost = repo.posts[len(repo.posts)-1]
	i, err := strconv.Atoi(latestPost.ID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(i + 1), nil
}

func (repo *MemoryBlogPostRepository) GetBlogs() []blog.Post {
	return repo.posts
}

func (repo *MemoryBlogPostRepository) PostBlog(post blog.Post) (bool, error) {
	if len(repo.posts) == 0 {
		repo.posts = []blog.Post{post}
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

func (repo *MemoryBlogPostRepository) LoadAllPosts(posts []blog.Post) {
	repo.posts = posts
}
