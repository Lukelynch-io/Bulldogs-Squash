package blog_claims

import "webservice/app/auth/claim"

const CREATE_BLOG claim.Claim = "CREATE_BLOG"

func GetBlogClaims() []claim.Claim {

	return []claim.Claim{
		CREATE_BLOG,
	}
}
