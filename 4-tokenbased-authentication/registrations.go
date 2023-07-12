package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func registerUser(username string, password string) (string, error) {
	db := dbConn()

	queryString := "INSERT INTO user(username, password) VALUES (?, ?)"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	_, err = stmt.Exec(username, hashedPassword)

	if err != nil {
		return "", err
	}

	return "Success\r\n", nil
}
