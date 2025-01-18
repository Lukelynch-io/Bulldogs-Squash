package post

import (
	"errors"
	"webservice/pkg/database"
)

type PostgresPostDatabase struct {
	database.PostgresDatabase
}

func (db *PostgresPostDatabase) GetPost(postId string) (Post, error) {
	sqlStatement := `
	SELECT * FROM "Posts"
	WHERE "Id" like $1;

	`
	var post Post
	rows, err := db.DB.Query(sqlStatement, postId)
	defer rows.Close()
	if err != nil {
		return post, err
	}

	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageUrl); err != nil {
			return post, err
		}
		return post, nil
	}
	return post, errors.New("No rows returned")
}

func (db *PostgresPostDatabase) GetPosts() ([]Post, error) {
	sqlStatement := `
	SELECT * FROM "Posts";
	`
	var posts []Post
	rows, err := db.DB.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return posts, err
	}

	for rows.Next() {
		var post Post

		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageUrl); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (db *PostgresPostDatabase) GetPostSnippets(trimAmount int) ([]Post, error) {
	sqlStatement := `
	SELECT Id, Title, substr(Description, 1, $1)
	FROM "Posts"
	`
	var posts []Post
	rows, err := db.DB.Query(sqlStatement, trimAmount)
	defer rows.Close()
	if err != nil {
		return posts, err
	}

	for rows.Next() {
		var post Post

		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageUrl); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
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
