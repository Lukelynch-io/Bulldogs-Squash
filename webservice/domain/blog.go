package domain

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type IBlogRepository interface {
	GetBlogs() []Post
	PostBlog(Post) (bool, error)
	LoadAllPosts([]Post)
}
