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
	hCon := controllers.NewHealthCheckController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", hCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres", gCon.GenresHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres/{genre_id}", gCon.GetGenreHandler).Methods(http.MethodGet)
	r.HandleFunc("/streams", sCon.PostStreamHandler).Methods(http.MethodPost)
	r.HandleFunc("/streams/{stream_id}", sCon.GetStreamHandler).Methods(http.MethodGet)

	r.Use(middlewares.PrepareResponse)

	return r
}
