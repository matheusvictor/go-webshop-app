package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connectDatabase() *sql.DB {
	connection := "user=postgres dbname=alura-webshop password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetDBInstance() *sql.DB {
	return connectDatabase()
}
