package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/repositories"
)

func GetGenreService() ([]models.Genres, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	genreList, err := repositories.SelectGenreList(db)
	if err != nil {
		return nil, err
	}

	return genreList, nil
}
