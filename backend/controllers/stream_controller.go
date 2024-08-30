package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sekibuuun/Devitch_ver2/backend/controllers/services"
	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

type StreamController struct {
	service services.StreamServicer
}

func NewStreamController(s services.StreamServicer) *StreamController {
	return &StreamController{service: s}
}

const (
	NoID       = iota
	MinGenreID = iota + 1
	MaxGenreID = iota + 10
)

func (c *StreamController) PostStreamHandler(w http.ResponseWriter, r *http.Request) {
	var streamRequest models.StreamRequest

	if err := json.NewDecoder(r.Body).Decode(&streamRequest); err != nil {
		log.Println("Could not decode request body to Go struct")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// genre_idsの検証
	if len(streamRequest.GenreIds) == 0 {
		log.Println("No genre_ids provided")
		http.Error(w, "Bad Request: At least one genre_id is required", http.StatusBadRequest)
		return
	}

	for _, genreId := range streamRequest.GenreIds {
		if genreId <= NoID || genreId > MaxGenreID {
			log.Println("Invalid genre_id:", genreId)
			http.Error(w, "Bad Request: Invalid genre_id", http.StatusBadRequest)
			return
		}
	}

	stream := models.Stream{
		Title:    streamRequest.Title,
		GenreIds: streamRequest.GenreIds,
	}

	newStream, err := c.service.PostStreamService(stream)
	if err != nil {
		log.Println("Could not post stream")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newStream); err != nil {
		log.Println("Could not encode stream to JSON")
	}
}
