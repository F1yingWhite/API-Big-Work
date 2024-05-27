package models

type History struct {
	CreatedAt string `json:"created_at"`
	MovieID   string `json:"movie_id"`
	UserID    string `json:"user_id"`
}
