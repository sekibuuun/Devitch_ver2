package repositories

import (
	"database/sql"
)

func InsertStreamGenre(db *sql.DB, stream_id int, genre_ids []int) error {
	const query = `INSERT INTO StreamGenre (stream_id, genre_id) VALUES (?, ?)`

	for _, genre_id := range genre_ids {
		if _, err := db.Exec(query, stream_id, genre_id); err != nil {
			return err
		}
	}

	return nil
}
