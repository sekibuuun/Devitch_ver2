package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/services/repositories"
)

type GenreServicer struct {
	repository repositories.GenreRepository
}

func NewGenreService(r repositories.GenreRepository) *GenreServicer {
	return &GenreServicer{repository: r}
}

func (s *GenreServicer) GetGenreService() ([]models.Genre, error) {
	genreList, err := s.repository.SelectGenreList()
	if err != nil {
		return nil, err
	}

	return genreList, nil
}

func (s *GenreServicer) GetGenreByIDService(genreID int) (models.Genre, error) {
	genre, err := s.repository.SelectGenre(genreID)
	if err != nil {
		return models.Genre{}, err
	}
	return genre, nil
}
