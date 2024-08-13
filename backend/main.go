package main

import (
	"log"
	"net/http"

	"github.com/sekibuuun/Devitch_ver2/backend/api"
)

func main() {
	// ルーターの初期化
	r := api.NewRouter()

	// サーバーの起動
	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
