package blog_test

import (
	"testing"
	"webservice/app/infra_interfaces"
	"webservice/infra"
)

func doSomethingwithIBlogRepo(repo infra_interfaces.IBlogRepository) {
	repo.GetBlogs()
}

func TestGetBlogPostsReturnsArrayOfString(t *testing.T) {
	blogRepo := &infra.MemoryBlogPostRepository{}
	doSomethingwithIBlogRepo(blogRepo)
}
