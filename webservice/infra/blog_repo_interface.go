package infra

type IBlogRepository interface {
	GetBlogs() []string
	PostBlog(string) (bool, error)
}
