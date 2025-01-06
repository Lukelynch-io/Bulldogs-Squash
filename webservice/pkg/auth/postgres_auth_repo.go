package auth

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresAuthDatabase struct {
	db *sql.DB
}

func (db *PostgresAuthDatabase) Connect(connStr string) {
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	db.db = database
}

func (db *PostgresAuthDatabase) SelectUserFromAuthDatabase(username string) *User {
	sqlStatement := `
	SELECT * from "Users"
	WHERE username like $1
	`
	rows, err := db.db.Query(sqlStatement, username)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	var foundUsers []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserId, &user.Username, &user.PasswordHash, &user.Role); err != nil {
			break
		}
		foundUsers = append(foundUsers, user)
	}
	if len(foundUsers) == 0 {
		panic("No Users found")
	}
	return &foundUsers[0]
}

// func (db *PostgresAuthDatabase) SelectFromPostgresTable() {
//
// 	rows, err := db.db.Query("SELECT * from Users")
// 	defer rows.Close()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var ingestedRows []Test
//
// 	for rows.Next() {
// 		var test Test
// 		if err := rows.Scan(&test.Id, &test.Name); err != nil {
// 			break
// 		}
// 		ingestedRows = append(ingestedRows, test)
// 	}
// 	fmt.Printf("rows: %v\n", ingestedRows)
// }

func (db *PostgresAuthDatabase) InsertTestUserIntoAuthDatabase(user User) {
	InserUserStatement := `
	INSERT INTO "Users" (userId, username, passwordHash, "role")
	VALUES ($1, $2, $3, $4)
	`
	_, err := db.db.Exec(InserUserStatement, user.UserId, user.Username, user.PasswordHash, user.Role)
	if err != nil {
		panic(err)
	}
}
