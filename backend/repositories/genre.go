package repositories

import (
	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

func (r *MyAppRepository) SelectGenreList() ([]models.Genre, error) {
	rows, err := r.db.Query("SELECT * FROM Genre")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []models.Genre
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.GenreId, &genre.Genre); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *MyAppRepository) SelectGenre(genreId int) (models.Genre, error) {
	const query = `SELECT * FROM Genre WHERE genre_id = ?`
	row := r.db.QueryRow(query, genreId)
	if err := row.Err(); err != nil {
		return models.Genre{}, err
	}

	var genre models.Genre
	err := row.Scan(&genre.GenreId, &genre.Genre)
	if err != nil {
		return models.Genre{}, err
	}

	return genre, nil
}
