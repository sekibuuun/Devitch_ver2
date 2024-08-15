package services

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	dbUser     = os.Getenv("MYSQL_ROOT_USER")
	dbPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
