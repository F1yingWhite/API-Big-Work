package models

import "time"

type History struct {
	CreatedAt *time.Time `json:"time"`
	MovieID   string     `json:"movie_id"`
	UserID    string     `json:"user_id"`
}
