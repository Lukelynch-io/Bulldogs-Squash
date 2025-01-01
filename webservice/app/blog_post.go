package app

import (
	"errors"
	"webservice/app/auth/claim"
	"webservice/app/blog/blog_claims"
)

func PostBlog(blogRepo IBlogRepository, post Post, actingUserClaims map[claim.Claim]claim.Claim) (bool, error) {
	if actingUserClaims[blog_claims.CREATE_BLOG] == blog_claims.CREATE_BLOG {
		blogRepo.PostBlog(post)
		return true, nil
	}
	return false, errors.New("User does not have permission to perform this action")
}
