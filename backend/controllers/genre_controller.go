package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sekibuuun/Devitch_ver2/backend/controllers/services"
	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

var hellos = []models.Hello{
	{ID: 1, Content: "Hello!"},
}

type GenreController struct {
	service services.GenreServicer
}

func NewGenreController(s services.GenreServicer) *GenreController {
	return &GenreController{service: s}
}

func (c *GenreController) HelloHandler(w http.ResponseWriter, r *http.Request) {
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
