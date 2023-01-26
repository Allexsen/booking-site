package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var msn string

func init() {
	user := os.Getenv("SQL_USER")
	pswd := os.Getenv("SQL_PSWD")
	adrs := os.Getenv("SQL_ADRS")
	dbnm := os.Getenv("SQL_DBNM")
	msn = user + ":" + pswd + "@tcp(" + adrs + ")/" + dbnm
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", msn)
	return db, err
}
