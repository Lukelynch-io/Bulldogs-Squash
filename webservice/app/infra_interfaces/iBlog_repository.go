package infra_interfaces

import "webservice/app/blog"

type IBlogRepository interface {
	GetBlogs() []blog.Post
	PostBlog(blog.Post) (bool, error)
	LoadAllPosts([]blog.Post)
}
