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

	genre_ids, err := InsertStreamGenre(db, int(stream_id), stream.GenreIds)

	if err != nil {
		return models.Stream{}, err
	}

	newStream.StreamId = int(stream_id)
	newStream.GenreIds = genre_ids

	return newStream, nil
}

func SelectStream(db *sql.DB, streamID int) (models.Stream, error) {
	const query = `SELECT s.stream_id, s.title, sg.genre_id FROM Stream s JOIN StreamGenre sg ON s.stream_id = sg.stream_id WHERE s.stream_id = ?;`

	row, err := db.Query(query, streamID)
	if err != nil {
		return models.Stream{}, err
	}
	defer row.Close()

	var stream models.Stream
	for row.Next() {
		var genre_id int
		if err := row.Scan(&stream.StreamId, &stream.Title, &genre_id); err != nil {
			return models.Stream{}, err
		}

		stream.GenreIds = append(stream.GenreIds, genre_id)
	}

	if err := row.Err(); err != nil {
		return models.Stream{}, err
	}

	return stream, nil

}
