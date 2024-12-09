package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type blogpost struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

// blogposts slice to seed record album data.
var blogposts = []blogpost{
	{
		ID:          "1",
		Title:       "2024 Summer Tournament",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velit!",
		ImageUrl:    "../../public/img/bulldogs.png",
	},
	{
		ID:          "2",
		Title:       "2023 Summer Tournament",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velit!",
		ImageUrl:    "../../public/img/bulldogs.png",
	},
	{
		ID:          "3",
		Title:       "2022 Summer Tournament",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velit!",
		ImageUrl:    "../../public/img/bulldogs.png",
	},
}

func getBlogPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogposts)
}

func main() {
	router := gin.Default()
	router.GET("/blogposts", getBlogPosts)

	router.Run("localhost:8080")
}
