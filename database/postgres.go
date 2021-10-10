package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "movietracker"
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}
