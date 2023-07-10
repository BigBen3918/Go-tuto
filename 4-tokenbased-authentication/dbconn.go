package main

import (
	"database/sql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"

	db, err := sql.Open(
		dbDriver,
		dbUser+":"+dbPass+"@/"+dbName,
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}
