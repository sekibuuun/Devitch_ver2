package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sekibuuun/Devitch_ver2/backend/api"
)

var (
	dbUser     = os.Getenv("MYSQL_ROOT_USER")
	dbPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("fail to connect db")
		return
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("failed to connect to DB %v", err)
	} else {
		fmt.Println("connect to DB")
	}
	// ルーターの初期化
	r := api.NewRouter()

	// サーバーの起動
	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
