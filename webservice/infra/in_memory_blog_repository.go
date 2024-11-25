package infra

type BlogPostRepository struct {
	posts []string
}

func (repo *BlogPostRepository) GetBlogs() []string {
	return repo.posts
}

func (repo *BlogPostRepository) PostBlog(post string) (bool, error) {
	if len(repo.posts) == 0 {
		repo.posts = []string{post}
		return true, nil
	}
	repo.posts = append(repo.posts, post)
	return true, nil
}
