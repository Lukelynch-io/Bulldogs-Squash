package post

import (
	"github.com/google/uuid"
)

// An In memory post storage that does not persist after program shutsdown
type InMemoryPostStorage struct {
	posts []Post
}

func (repo *InMemoryPostStorage) GetPosts() ([]Post, error) {
	return repo.posts, nil
}

func (repo *InMemoryPostStorage) GetPostSnippets() ([]Post, error) {
	var trimmedPosts []Post
	for _, post := range repo.posts {
		if len(post.Description) > 100 {
			trimmedDescription := post.Description[:100]
			tempPost := Post{
				post.ID,
				post.Title,
				trimmedDescription,
				post.ImageUrl,
			}
			trimmedPosts = append(trimmedPosts, tempPost)
			continue
		}
		trimmedPosts = append(trimmedPosts, post)
	}
	return trimmedPosts, nil
}

func (repo *InMemoryPostStorage) InsertPost(post Post) (bool, error) {
	if post.ID == "" {
		post.ID = uuid.NewString()
	}
	repo.posts = append(repo.posts, post)
	return true, nil
}

func (repo *InMemoryPostStorage) LoadAllPosts(posts []Post) {
	repo.posts = posts
}
