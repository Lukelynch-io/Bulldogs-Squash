package post

import (
	"errors"
	"mime/multipart"
	"webservice/pkg/auth"
	"webservice/pkg/filestore"
)

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type NewPost struct {
	PostTitle       string                `form:"postTitle" binding:"required"`
	PostDescription string                `form:"postDescription" binding:"required"`
	FileData        *multipart.FileHeader `form:"imageFile" binding:"required"`
}

type PostStorage interface {
	GetPosts() ([]Post, error)
	GetPostSnippets() ([]Post, error)
	InsertPost(Post) (bool, error)
	LoadAllPosts([]Post)
}

func GetPosts(repo PostStorage, trimDescriptions bool) ([]Post, error) {
	if trimDescriptions {
		return repo.GetPosts()
	}
	posts, err := repo.GetPostSnippets()
	if err != nil {
		return posts, err
	}
	for i := 0; i < len(posts); i++ {
		posts[i].Description = posts[i].Description + "..."
	}
	return posts, nil
}

func InsertPost(blogRepo PostStorage, fileStorage filestore.Filestore, post NewPost, actingUserClaims map[auth.Claim]auth.Claim) (bool, error) {
	if actingUserClaims[auth.CREATE_BLOG] != auth.CREATE_BLOG {
		return false, errors.New("User does not have permission to perform this action")
	}
	filepath, err := fileStorage.StoreFile(post.FileData)
	if err != nil {
		return false, errors.New("Failed to upload image")
	}
	parsedPost := Post{
		Title:       post.PostTitle,
		Description: post.PostDescription,
		ImageUrl:    filepath,
	}
	blogRepo.InsertPost(parsedPost)
	return true, nil
}
