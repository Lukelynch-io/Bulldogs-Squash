package main

import (
	"os"
	"webservice/app/infra_interfaces"
	"webservice/infra"
	"webservice/routes/auth"
	"webservice/routes/blogpost"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var iblogPostRepository infra_interfaces.IBlogRepository

func main() {
	var secretKey = os.Getenv("JWT_Secret_Key")
	if secretKey == "" {
		os.Exit(1)
	}
	iblogPostRepository = &infra.MemoryBlogPostRepository{}
	iblogPostRepository.LoadAllPosts(testdata.LoadDummyData())

	router := gin.Default()
	auth.Routes(router, []byte(secretKey))
	blogpost.Routes(router, iblogPostRepository)

	router.RunTLS("localhost:8080", "./certs/dev-cert.pem", "./certs/dev-key.pem")
}
