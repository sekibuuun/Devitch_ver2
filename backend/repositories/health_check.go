package repositories

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

func SelectHelloList() ([]models.Hello, error) {
	hellos := []models.Hello{
		{ID: 1, Content: "Hello!"},
	}

	return hellos, nil
}
