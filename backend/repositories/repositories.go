package repositories

import (
	"database/sql"
)

type MyAppRepository struct {
	db *sql.DB
}

func NewMyAppRepository(db *sql.DB) *MyAppRepository {
	return &MyAppRepository{db: db}
}
