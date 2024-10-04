package repositories

import (
	"database/sql"

	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

type HealthCheckRepository struct {
	db *sql.DB
}

func NewHealthCheckRepository(db *sql.DB) *HealthCheckRepository {
	return &HealthCheckRepository{db: db}
}

func (r *HealthCheckRepository) SelectHelloList() ([]models.Hello, error) {
	hellos := []models.Hello{
		{ID: 1, Content: "Hello!"},
	}

	return hellos, nil
}
