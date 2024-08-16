package models

type Hello struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Genre struct {
	GenreId int    `json:"genre_id"`
	Genre   string `json:"genre"`
}

type Stream struct {
	StreamId        int       `json:"stream_id"`
	Title           string    `json:"title"`
	GenreIds        []int     `json:"genre_ids"`
	Comments        []Comment `json:"comments"`
	CommentsLength  int       `json:"comment_length"`
	Listeners       []User    `json:"listeners"`
	ListenersLength int       `json:"listener_length"`
}

type Comment struct {
	CommentId int    `json:"comment_id"`
	StreamId  int    `json:"stream_id"`
	UserId    int    `json:"user_id"`
	Content   string `json:"content"`
	SendAt    string `json:"send_at"`
}

type User struct {
	UserID   int      `json:"user_id"`
	UserName string   `json:"user_name"`
	Links    []string `json:"links"`
}

type StreamGenre struct {
	StreamId int   `json:"stream_id"`
	GenreId  []int `json:"genre_id"`
}

type StreamRequest struct {
	Title    string `json:"title"`
	GenreIds []int  `json:"genre_ids"`
}
