package services

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
	"github.com/sekibuuun/Devitch_ver2/backend/repositories"
)

func (s *MyAppService) GetHelloService() ([]models.Hello, error) {
	hello, err := repositories.SelectHelloList()
	if err != nil {
		return nil, err
	}
	return hello, nil
}
