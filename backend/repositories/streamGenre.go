package repositories

import (
	"database/sql"
)

func InsertStreamGenre(db *sql.DB, stream_id int, genre_ids []int) ([]int, error) {
	const query = `INSERT INTO StreamGenre (stream_id, genre_id) VALUES (?, ?)`

	genres := make([]int, 0, len(genre_ids))

	for _, genre_id := range genre_ids {
		if _, err := db.Exec(query, stream_id, genre_id); err != nil {
			return nil, err
		}
		genres = append(genres, genre_id)
	}

	return genres, nil
}
