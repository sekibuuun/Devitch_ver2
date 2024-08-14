package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sekibuuun/Devitch_ver2/backend/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres", handlers.GenresHandler).Methods(http.MethodGet)
	return r
}
