package blog_test

import (
	"testing"
	"webservice/infra"
)

func doSomethingwithIBlogRepo(repo infra.IBlogRepository) {
	repo.GetBlogs()
}

func TestGetBlogPostsReturnsArrayOfString(t *testing.T) {
	blogRepo := new(infra.BlogPostRepository)
	doSomethingwithIBlogRepo(blogRepo)

}
