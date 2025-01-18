package post

import (
	"errors"

	"github.com/google/uuid"
)

// An In memory post storage that does not persist after program shutsdown
type InMemoryPostStorage struct {
	posts []Post
}

func (repo *InMemoryPostStorage) GetPost(postId string) (Post, error) {
	for _, post := range repo.posts {
		if post.ID == postId {
			return post, nil
		}
	}
	var noPost Post
	return noPost, errors.New("Could not find post with matching ID")
}

func (repo *InMemoryPostStorage) GetPosts() ([]Post, error) {
	return repo.posts, nil
}

func (repo *InMemoryPostStorage) GetPostSnippets(trimAmount int) ([]Post, error) {
	var trimmedPosts []Post
	for _, post := range repo.posts {
		if len(post.Description) > trimAmount {
			trimmedDescription := post.Description[:trimAmount]
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
