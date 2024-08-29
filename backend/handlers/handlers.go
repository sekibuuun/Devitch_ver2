package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/services"
)

const (
	NoID       = iota
	MinGenreID = iota + 1
	MaxGenreID = iota + 10
)

var hellos = []models.Hello{
	{ID: 1, Content: "Hello!"},
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&hellos); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}

func GenresHandler(w http.ResponseWriter, r *http.Request) {
	genres, err := services.GetGenreService()
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

func PostStreamHandler(w http.ResponseWriter, r *http.Request) {
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

	newStream, err := services.PostStreamService(stream)
	if err != nil {
		log.Println("Could not post stream")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newStream); err != nil {
		log.Println("Could not encode stream to JSON")
	}
}
