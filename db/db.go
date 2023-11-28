package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB //type *sql.DB?

func SetupDatabase() {
	var err error //type?

	DB, err = sql.Open("postgres", "user=postgres password=mypassword host=localhost port=5432 dbname=movies sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
