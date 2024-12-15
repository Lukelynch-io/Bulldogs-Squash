package blog

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

func AddBlogPost(postArray []Post, newPost Post) []Post {
	postArray = append(postArray, newPost)
	return postArray
}
