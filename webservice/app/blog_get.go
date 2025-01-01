package app

import "webservice/domain"

func GetPosts(repo domain.IBlogRepository) []domain.Post {
	return repo.GetBlogs()
}
