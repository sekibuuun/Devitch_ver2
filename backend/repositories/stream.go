package repositories

import (
	"database/sql"

	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

func InsertStream(db *sql.DB, stream models.Stream) (models.Stream, error) {
	const query = `INSERT INTO Stream (title) VALUES (?)`

	var newStream models.Stream
	newStream.Title = stream.Title

	result, err := db.Exec(query, newStream.Title)
	if err != nil {
		return models.Stream{}, err
	}

	stream_id, _ := result.LastInsertId()

	if err := InsertStreamGenre(db, int(stream_id), stream.GenreIds); err != nil {
		return models.Stream{}, err
	}

	newStream.StreamId = int(stream_id)

	return newStream, nil
}
