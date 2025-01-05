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

var IBlogPostRepository domain.IBlogRepository = new(infra.MemoryBlogPostRepository)
var IAuthUserRepository domain.IAuthRepo
var SecretKey []byte

const JWT_SECRET_KEY = "Jwt_Secret_Key"

func setupSecretKey(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(env.SecretKey, secretKey)
	}
}

func setupAuthRepo(c *gin.Context) {
	c.Set(env.AuthRepo, IAuthUserRepository)
}

func setupBlogRepo(c *gin.Context) {
	c.Set(env.BlogRepo, IBlogPostRepository)
}

func init() {
	authRepo := infra.NewMemoryAuthRepository()
	IAuthUserRepository = &authRepo
	secretKey := os.Getenv(JWT_SECRET_KEY)
	if secretKey == "" {
		log.Fatal(fmt.Sprintf("%s was not defined", JWT_SECRET_KEY))
	}
	SecretKey = []byte(secretKey)
	IBlogPostRepository.LoadAllPosts(testdata.LoadDummyData())
}

func main() {
	router := gin.New()
	router.Use(
		setupSecretKey(SecretKey),
		setupAuthRepo,
		setupBlogRepo)
	routes.LoadAuthRoutes(router)
	routes.LoadBlogPostRoutes(router)

	router.RunTLS("localhost:8080", "./certs/dev-cert.pem", "./certs/dev-key.pem")
}
