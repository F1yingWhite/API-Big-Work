package models

import "time"

type History struct {
	CreatedAt *time.Time `json:"time"`
	MovieID   int        `json:"movie_id"`
	UserID    string     `json:"user_id"`
}

func CreateHistory(movieID int, userID string) error {
	history := History{
		CreatedAt: nil,
		MovieID:   movieID,
		UserID:    userID,
	}
	return DB.Create(&history).Error
}
