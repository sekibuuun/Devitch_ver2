package repositories

import (
	"database/sql"

	"github.com/sekibuuun/Devitch_ver2/backend/models"
)

type StreamRepository struct {
	db              *sql.DB
	streamGenreRepo *StreamGenreRepository
}

func NewStreamRepository(db *sql.DB) *StreamRepository {
	return &StreamRepository{
		db:              db,
		streamGenreRepo: NewStreamGenreRepository(db),
	}
}

func (r *StreamRepository) InsertStream(stream models.Stream) (models.Stream, error) {
	const query = `INSERT INTO Stream (title) VALUES (?)`

	var newStream models.Stream
	newStream.Title = stream.Title

	result, err := r.db.Exec(query, newStream.Title)
	if err != nil {
		return models.Stream{}, err
	}

	stream_id, _ := result.LastInsertId()

	genre_ids, err := r.streamGenreRepo.InsertStreamGenre(int(stream_id), stream.GenreIds)
	if err != nil {
		return models.Stream{}, err
	}

	newStream.StreamId = int(stream_id)
	newStream.GenreIds = genre_ids

	return newStream, nil
}

func (r *StreamRepository) SelectStream(streamID int) (models.Stream, error) {
	const query = `SELECT s.stream_id, s.title, sg.genre_id FROM Stream s JOIN StreamGenre sg ON s.stream_id = sg.stream_id WHERE s.stream_id = ?;`

	row, err := r.db.Query(query, streamID)
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
