package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sekibuuun/Devitch_ver2/backend/api/middlewares"
	"github.com/sekibuuun/Devitch_ver2/backend/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres", handlers.GenresHandler).Methods(http.MethodGet)
	r.HandleFunc("/streams", handlers.PostStreamHandler).Methods(http.MethodPost)

	r.Use(middlewares.CORS)
	r.Use(middlewares.JSON)

	return r
}
