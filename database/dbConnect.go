package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// dbPassword := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/go_tweet")
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
