package db

import (
	"database/sql"
	"log"
)

func ConnectionDatabasePostgres() *sql.DB {
	connStr := "user=postgres dbname=postgres password=alura host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
