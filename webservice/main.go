package main

import (
	"webservice/app/infra_interfaces"
	"webservice/infra"
	"webservice/routes/auth"
	"webservice/routes/blogpost"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var iblogPostRepository infra_interfaces.IBlogRepository

func main() {
	iblogPostRepository = &infra.MemoryBlogPostRepository{}
	iblogPostRepository.LoadAllPosts(testdata.LoadDummyData())

	router := gin.Default()
	auth.Routes(router)
	blogpost.Routes(router, iblogPostRepository)

	router.Run("localhost:8080")
}
