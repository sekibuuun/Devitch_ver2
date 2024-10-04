package repositories

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

type HealthCheckRepository interface {
	SelectHelloList() ([]models.Hello, error)
}
type GenreRepository interface {
	SelectGenreList() ([]models.Genre, error)
	SelectGenre(genreId int) (models.Genre, error)
}

type StreamRepository interface {
	InsertStream(models.Stream) (models.Stream, error)
	SelectStream(int) (models.Stream, error)
}

type StreamGenreRepository interface {
	InsertStreamGenre(int, []int) ([]int, error)
}
