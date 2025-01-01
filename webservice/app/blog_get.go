package app

func GetPosts(repo IBlogRepository) []Post {
	return repo.GetBlogs()
}
