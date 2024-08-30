package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sekibuuun/Devitch_ver2/backend/api/middlewares"
	"github.com/sekibuuun/Devitch_ver2/backend/controllers"
	"github.com/sekibuuun/Devitch_ver2/backend/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	gCon := controllers.NewGenreController(ser)
	sCon := controllers.NewStreamController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", gCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres", gCon.GenresHandler).Methods(http.MethodGet)
	r.HandleFunc("/streams", sCon.PostStreamHandler).Methods(http.MethodPost)

	r.Use(middlewares.CORS)
	r.Use(middlewares.JSON)

	return r
}
