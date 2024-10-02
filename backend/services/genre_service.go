package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/repositories"
)

func (s *MyAppService) GetGenreService() ([]models.Genre, error) {
	genreList, err := repositories.SelectGenreList(s.db)
	if err != nil {
		return nil, err
	}

	return genreList, nil
}

func (s *MyAppService) GetGenreByIDService(genreID int) (models.Genre, error) {
	genre, err := repositories.SelectGenre(s.db, genreID)
	if err != nil {
		return models.Genre{}, err
	}
	return genre, nil
}
