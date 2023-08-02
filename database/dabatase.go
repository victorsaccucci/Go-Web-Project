package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBUser     = "root"
	DBPassword = "gowebprojectdb"
	DBName     = "admin"
)

func OpenDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", DBUser, DBPassword, DBName)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
