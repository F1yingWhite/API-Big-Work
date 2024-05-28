package models

import "time"

type History struct {
	CreatedAt *time.Time `json:"time"`
	MovieID   uint       `json:"movie_id"`
	UserID    string     `json:"user_id"`
}

func CreateHistory(movieID uint, userID string) error {
	now := time.Now()
	history := History{
		MovieID:   movieID,
		UserID:    userID,
		CreatedAt: &now, //记录创建时间 like相同则推荐最新的
	}
	return DB.Create(&history).Error
}
