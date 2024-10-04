package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/services/repositories"
)

type HealthCheckServicer struct {
	repository repositories.HealthCheckRepository
}

func NewHealthCheckService(r repositories.HealthCheckRepository) *HealthCheckServicer {
	return &HealthCheckServicer{repository: r}
}

func (s *HealthCheckServicer) GetHelloService() ([]models.Hello, error) {
	hello, err := s.repository.SelectHelloList()
	if err != nil {
		return nil, err
	}
	return hello, nil
}
