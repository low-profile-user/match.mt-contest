package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectToDatabase() *sql.DB {
	connectionConfig := "user=root dbname=matchcontest password=postgres sslmode=disable host=db"
	db, error := sql.Open("postgres", connectionConfig)

	if error != nil {
		panic(error.Error())
	}

	return db
}
