package infra

import "webservice/app/blog"

type MemoryBlogPostRepository struct {
	posts []blog.Post
}

func (repo *MemoryBlogPostRepository) GetBlogs() []blog.Post {
	return repo.posts
}

func (repo *MemoryBlogPostRepository) PostBlog(post blog.Post) (bool, error) {
	if len(repo.posts) == 0 {
		repo.posts = []blog.Post{post}
		return true, nil
	}
	repo.posts = append(repo.posts, post)
	return true, nil
}

func (repo *MemoryBlogPostRepository) LoadAllPosts(posts []blog.Post) {
	repo.posts = posts
}
