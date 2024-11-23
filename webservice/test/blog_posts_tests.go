package tests

import (
	"testing"
	"webservice/blog"
)

func getBlogPostsReturnsArrayOfString(t *testing.T) {
	blog.GetBlogPosts()

}
