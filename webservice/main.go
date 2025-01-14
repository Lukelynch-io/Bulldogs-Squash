package main

import (
	"fmt"
	"log"
	"os"
	"webservice/env"
	"webservice/pkg/auth"
	"webservice/pkg/post"
	"webservice/routes"
	"webservice/testdata"

	"github.com/gin-gonic/gin"
)

var IBlogPostRepository post.PostStorage = new(post.InMemoryPostStorage)
var IAuthUserRepository auth.UserDataStorage
var TokenStorage auth.TokenStorage
var SecretKey []byte

const JWT_SECRET_KEY = "Jwt_Secret_Key"

func setupSecretKey(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(env.SecretKey, secretKey)
	}
}

func setupAuthRepo(c *gin.Context) {
	c.Set(env.AuthRepo, IAuthUserRepository)
	c.Set(env.TokenStorage, TokenStorage)
}

func setupBlogRepo(c *gin.Context) {
	c.Set(env.PostRepo, IBlogPostRepository)
}

func init() {
	authRepo := auth.NewInMemoryAuthRepository()
	IAuthUserRepository = &authRepo
	TokenStorage = &authRepo
	secretKey := os.Getenv(JWT_SECRET_KEY)
	if secretKey == "" {
		log.Fatal(fmt.Sprintf("%s was not defined", JWT_SECRET_KEY))
	}
	SecretKey = []byte(secretKey)
	IBlogPostRepository.LoadAllPosts(testdata.LoadDummyData())
}

func main() {
	gin.ForceConsoleColor()
	router := gin.New()

	router.Use(
		setupSecretKey(SecretKey),
		setupAuthRepo,
		setupBlogRepo)
	routes.LoadAuthRoutes(router)
	routes.LoadUserRoutes(router)
	routes.LoadBlogPostRoutes(router)

	router.RunTLS("localhost:8080", "./certs/dev-cert.pem", "./certs/dev-key.pem")
}
