package database

import (
	"database/sql"
	"log"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func (db *PostgresDatabase) Connect(connStr string) {
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	db.DB = database
}
