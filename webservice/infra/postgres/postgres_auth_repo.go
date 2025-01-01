package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Test struct {
	Id   int32
	Name string
}

func SelectFromPostgresTable(connStr string) {
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := database.Query("SELECT * from testdb")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var ingestedRows []Test

	for rows.Next() {
		var test Test
		if err := rows.Scan(&test.Id, &test.Name); err != nil {
			break
		}
		ingestedRows = append(ingestedRows, test)
	}
	fmt.Printf("rows: %v\n", ingestedRows)
}
