package models

type Hello struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Genres struct {
	GenreId int    `json:"genre_id"`
	Genre   string `json:"genre"`
}
