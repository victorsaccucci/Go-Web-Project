package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var ErrNoRows = errors.New("no rows found")

const (
	DBUser     = "root"
	DBPassword = "admin"
	DBName     = "gowebprojectdb"
	DBPort     = "3306"
)

func OpenDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", DBUser, DBPassword, DBPort, DBName)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
		panic(err.Error())
	}
	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
