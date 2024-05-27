package models

import "time"

type History struct {
	CreatedAt *time.Time `json:"time"`
	MovieID   uint        `json:"movie_id"`
	UserID    string     `json:"user_id"`
}

func CreateHistory(movieID uint, userID string) error {
	history := History{
		MovieID: movieID,
		UserID:  userID,
	}
	return DB.Create(&history).Error
}
