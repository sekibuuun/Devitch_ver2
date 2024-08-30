package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

type GenreServicer interface {
	GetGenreService() ([]models.Genre, error)
}

type StreamServicer interface {
	PostStreamService(stream models.Stream) (models.Stream, error)
}
