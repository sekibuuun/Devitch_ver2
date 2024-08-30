package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/repositories"
)

func (s *MyAppService) PostStreamService(stream models.Stream) (models.Stream, error) {
	newStream, err := repositories.InsertStream(s.db, stream)
	if err != nil {
		return models.Stream{}, err
	}

	return newStream, nil
}
