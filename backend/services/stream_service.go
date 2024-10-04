package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/services/repositories"
)

type StreamServicer struct {
	repository repositories.StreamRepository
}

func NewStreamService(r repositories.StreamRepository) *StreamServicer {
	return &StreamServicer{repository: r}
}

func (s *StreamServicer) PostStreamService(stream models.Stream) (models.Stream, error) {
	newStream, err := s.repository.InsertStream(stream)
	if err != nil {
		return models.Stream{}, err
	}

	return newStream, nil
}

func (s *StreamServicer) GetStreamService(streamID int) (models.Stream, error) {
	stream, err := s.repository.SelectStream(streamID)
	if err != nil {
		return models.Stream{}, err
	}

	return stream, nil
}
