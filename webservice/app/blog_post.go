package app

import (
	"errors"
	"webservice/domain"
	"webservice/domain/blog_claims"
)

func PostBlog(blogRepo domain.IBlogRepository, post domain.Post, actingUserClaims map[domain.Claim]domain.Claim) (bool, error) {
	if actingUserClaims[blog_claims.CREATE_BLOG] == blog_claims.CREATE_BLOG {
		blogRepo.PostBlog(post)
		return true, nil
	}
	return false, errors.New("User does not have permission to perform this action")
}
