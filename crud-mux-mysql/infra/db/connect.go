package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
	driver := "mysql"
	user := "root"
	pass := ""
	database := "go-crud"

	db, err := sql.Open(driver, user+":"+pass+"@/"+database)
	if err != nil {
		panic(err.Error())
	}

	return db
}
