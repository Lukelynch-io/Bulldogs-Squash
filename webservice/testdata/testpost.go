package testdata

import "webservice/pkg/post"

// blogposts slice to seed record album data.
var blogposts = []post.Post{
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
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velit!Lorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velitLorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velitLorem ipsum dolor sit amet, consectetur adipisicing elit. Qui dolorum reiciendis itaque, vitae mollitia ad accusantium similique pariatur minus repellat maiores voluptatibus recusandae perspiciatis fugiat provident nostrum assumenda inventore velit",
		ImageUrl:    "../../public/img/bulldogs.png",
	},
}

func LoadDummyData() []post.Post {
	return blogposts
}
