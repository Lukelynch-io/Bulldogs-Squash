package testdata

import "webservice/domain"

// blogposts slice to seed record album data.
var blogposts = []domain.Post{
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

func LoadDummyData() []domain.Post {
	return blogposts
}
