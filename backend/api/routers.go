package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sekibuuun/Devitch_ver2/backend/api/middlewares"
	"github.com/sekibuuun/Devitch_ver2/backend/controllers"
	"github.com/sekibuuun/Devitch_ver2/backend/repositories"
	"github.com/sekibuuun/Devitch_ver2/backend/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	hRep := repositories.NewHealthCheckRepository(db)
	gRep := repositories.NewGenreRepository(db)
	sRep := repositories.NewStreamRepository(db)

	hSer := services.NewHealthCheckService(hRep)
	gSer := services.NewGenreService(gRep)
	sSer := services.NewStreamService(sRep)

	hCon := controllers.NewHealthCheckController(hSer)
	gCon := controllers.NewGenreController(gSer)
	sCon := controllers.NewStreamController(sSer)

	r := mux.NewRouter()
	r.HandleFunc("/hello", hCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres", gCon.GenresHandler).Methods(http.MethodGet)
	r.HandleFunc("/genres/{genre_id}", gCon.GetGenreHandler).Methods(http.MethodGet)
	r.HandleFunc("/streams", sCon.PostStreamHandler).Methods(http.MethodPost)
	r.HandleFunc("/streams/{stream_id}", sCon.GetStreamHandler).Methods(http.MethodGet)

	r.Use(middlewares.PrepareResponse)

	return r
}
