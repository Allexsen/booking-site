package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Connect() *sql.DB {
	user := os.Getenv("SQL_USER")
	pswd := os.Getenv("SQL_PSWD")
	adrs := os.Getenv("SQL_ADRS")
	dbnm := os.Getenv("SQL_DBNM")
	msn := user + ":" + pswd + "@tcp(" + adrs + ")/" + dbnm

	d, err := sql.Open("mysql", msn)
	if err != nil {
		panic(err)
	}

	db = d
	return db
}

func GetDB() *sql.DB { // return DB instance
	return db
}
