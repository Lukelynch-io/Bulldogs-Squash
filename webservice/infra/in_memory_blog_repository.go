package infra

type MemoryBlogPostRepository struct {
	posts []string
}

func (repo *MemoryBlogPostRepository) GetBlogs() []string {
	return repo.posts
}

func (repo *MemoryBlogPostRepository) PostBlog(post string) (bool, error) {
	if len(repo.posts) == 0 {
		repo.posts = []string{post}
		return true, nil
	}
	repo.posts = append(repo.posts, post)
	return true, nil
}
