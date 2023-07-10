package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func generateToken(username string, password string) (map[string]interface{}, error) {
	db := dbConn()

	queryString := "SELECT username, password FROM user WHERE username=?"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	userId := 0
	accountPassword := ""

	err = stmt.QueryRow(username).Scan(&userId, &accountPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid username or password.\r\n")
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(password))

	if err != nil {
		return nil, errors.New("Invalid username or password.\r\n")
	}

	queryString = "INSERT INTO authtication_tokens(user_id, auth_token, generated_at, expired_at) VALUES(?, ?, ?, ?)"

	stmt, err = db.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	randomToken := make([]byte, 32)

	_, err = rand.Read(randomToken)

	if err != nil {
		return nil, err
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)

	const timeLayout = "2006-01-02 15:04:05"

	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 60)

	generatedAt := dt.Format(timeLayout)
	expiredAt := expirtyTime.Format(timeLayout)

	_, err = stmt.Exec(userId, authToken, generatedAt, expiredAt)

	if err != nil {
		return nil, err
	}

	tokenDetails := map[string]interface{}{
		"token_type":   "Bearer",
		"auth_token":   authToken,
		"generated_at": generatedAt,
		"expired_at":   expiredAt,
	}

	return tokenDetails, nil
}

func validateToken(authToken string) (map[string]interface{}, error) {
	db := dbConn()

	queryString := `SELECT id, username, generated_at, expires_at FROM token LEFT JOIN user on token.id=?`

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	userId := 0
	username := ""
	generatedAt := ""
	expiresAt := ""

	err = stmt.QueryRow(authToken).Scan(&userId, &username, &generatedAt, &expiresAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid access token.\r\n")
		}

		return nil, err
	}

	const timeLayout = "2006-01-02 15:04:05"

	expiryTime, _ := time.Parse(timeLayout, expiresAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return nil, errors.New("The  token is expired. \r\n")
	}

	userDetails := map[string]interface{}{
		"user_id":      userId,
		"username":     username,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}

	return userDetails, nil
}
