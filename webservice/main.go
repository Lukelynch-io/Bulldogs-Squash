package main

import (
	"fmt"
	"log"
	"os"
	"webservice/domain"
	"webservice/env"
	"webservice/infra"
	"webservice/routes"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var IBlogPostRepository domain.IBlogRepository
var IAuthUserRepository domain.IAuthRepo
var SecretKey []byte

const JWT_SECRET_KEY = "Jwt_Secret_Key"

func setupSecretKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(env.SecretKey, SecretKey)
	}
}

func setupAuthRepo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(env.AuthRepo, IAuthUserRepository)
	}
}

func setupBlogRepo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(env.BlogRepo, IBlogPostRepository)
	}
}

func setupDependencies() {
	IAuthUserRepository = infra.NewMemoryAuthRepository()
	IBlogPostRepository = new(infra.MemoryBlogPostRepository)
	var secretKey = os.Getenv(JWT_SECRET_KEY)
	if secretKey == "" {
		log.Fatal(fmt.Sprintf("%s was not defined", JWT_SECRET_KEY))
	}
	SecretKey = []byte(secretKey)
	IBlogPostRepository.LoadAllPosts(testdata.LoadDummyData())
}

func main() {
	setupDependencies()

	router := gin.New()
	router.Use(
		setupSecretKey(),
		setupAuthRepo(),
		setupBlogRepo())
	routes.LoadAuthRoutes(router)
	routes.LoadBlogPostRoutes(router)

	router.RunTLS("localhost:8080", "./certs/dev-cert.pem", "./certs/dev-key.pem")
}
