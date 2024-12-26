package main

import (
	"fmt"
	"log"
	"os"
	"webservice/app/blog"
	"webservice/infra"
	"webservice/routes/auth"
	"webservice/routes/blogpost"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var iblogPostRepository blog.IBlogRepository

const JWT_SECRET_KEY = "Jwt_Secret_Key"

func main() {
	var secretKey = os.Getenv(JWT_SECRET_KEY)
	if secretKey == "" {
		log.Fatal(fmt.Sprintf("%s was not defined", JWT_SECRET_KEY))
	}
	iblogPostRepository = &infra.MemoryBlogPostRepository{}
	iblogPostRepository.LoadAllPosts(testdata.LoadDummyData())

	router := gin.Default()
	auth.Routes(router, []byte(secretKey))
	blogpost.Routes(router, iblogPostRepository)

	router.RunTLS("localhost:8080", "./certs/dev-cert.pem", "./certs/dev-key.pem")
}
