package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sekibuuun/Devitch_ver2/backend/controllers/services"
)

type GenreController struct {
	service services.GenreServicer
}

func NewGenreController(s services.GenreServicer) *GenreController {
	return &GenreController{service: s}
}

func (c *GenreController) GenresHandler(w http.ResponseWriter, r *http.Request) {
	genres, err := c.service.GetGenreService()
	if err != nil {
		log.Println("Could not get genres")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&genres); err != nil {
		log.Println("Could not encode genres to JSON")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println(buf.String())

	_, err = fmt.Fprint(w, buf.String())
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func (c *GenreController) GetGenreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genreID := vars["genre_id"]

	log.Println("genre_id:", genreID)

	id, err := strconv.Atoi(genreID)
	if err != nil {
		log.Println("Invalid genre ID")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	genre, err := c.service.GetGenreByIDService(id)
	if err != nil {
		log.Println("Could not get genre")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&genre); err != nil {
		log.Println("Could not encode genre to JSON")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println(buf.String())

	_, err = fmt.Fprint(w, buf.String())
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
