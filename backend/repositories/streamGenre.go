package repositories

import (
	"strings"
)

func (r *MyAppRepository) InsertStreamGenre(stream_id int, genre_ids []int) ([]int, error) {
	const query = `INSERT INTO StreamGenre (stream_id, genre_id) VALUES `

	placeholders := make([]string, len(genre_ids))
	args := make([]interface{}, 0, len(genre_ids)*2)

	for i, genre_id := range genre_ids {
		placeholders[i] = "(?, ?)"
		args = append(args, stream_id, genre_id)
	}

	fullQuery := query + strings.Join(placeholders, ", ")

	_, err := r.db.Exec(fullQuery, args...)
	if err != nil {
		return nil, err
	}

	return genre_ids, nil
}
