package repositories

import (
	"database/sql"

	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

func SelectGenreList(db *sql.DB) ([]models.Genres, error) {
	rows, err := db.Query("SELECT * FROM Genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []models.Genres
	for rows.Next() {
		var genre models.Genres
		if err := rows.Scan(&genre.GenreId, &genre.Genre); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}
