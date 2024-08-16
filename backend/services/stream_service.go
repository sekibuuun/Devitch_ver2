package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/repositories"
)

func PostStreamService(stream models.Stream) (models.Stream, error) {
	db, err := connectDB()
	if err != nil {
		return models.Stream{}, err
	}
	defer db.Close()

	newStream, err := repositories.InsertStream(db, stream)
	if err != nil {
		return models.Stream{}, err
	}

	return newStream, nil
}
