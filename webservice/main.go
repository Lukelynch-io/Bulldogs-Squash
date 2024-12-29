package main

import (
	"fmt"
	"log"
	"os"
	"webservice/app/auth"
	"webservice/app/blog"
	"webservice/infra"
	auth_route "webservice/routes/auth"
	"webservice/routes/blogpost"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var iblogPostRepository blog.IBlogRepository
var iAuthUserRepository auth.IAuthRepo

const JWT_SECRET_KEY = "Jwt_Secret_Key"

func main() {
	var secretKey = os.Getenv(JWT_SECRET_KEY)
	if secretKey == "" {
		log.Fatal(fmt.Sprintf("%s was not defined", JWT_SECRET_KEY))
	}
	iblogPostRepository = new(infra.MemoryBlogPostRepository)
	iblogPostRepository.LoadAllPosts(testdata.LoadDummyData())
	iAuthUserRepository = new(infra.MemoryAuthRepository)

	router := gin.Default()
	auth_route.Routes(router, []byte(secretKey), iAuthUserRepository)
	blogpost.Routes(router, iblogPostRepository)

	router.RunTLS("localhost:8080", "./certs/dev-cert.pem", "./certs/dev-key.pem")
}
