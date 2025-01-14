package post

import "webservice/pkg/database"

type PostgresPostDatabase struct {
	database.PostgresDatabase
}

func (db *PostgresPostDatabase) GetPosts() []Post {
	sqlStatement := `
	SELECT * FROM "Posts";
	`
	var posts []Post
	rows, err := db.DB.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var post Post

		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageUrl); err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return posts
}

func (db *PostgresPostDatabase) InsertPost(post Post) (bool, error) {
	InsertPostStatement := `
	INSERT INTO "Posts" (Id, Title, Description, ImageUrl)
	VALUES ($1, $2, $3, $4)
	`
	_, err := db.DB.Exec(InsertPostStatement, post.ID, post.Title, post.Description, post.ImageUrl)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (db *PostgresPostDatabase) LoadAllPosts(posts []Post) {
	for _, post := range posts {
		db.InsertPost(post)
	}
}
