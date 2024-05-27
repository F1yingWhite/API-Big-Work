package models

type Like struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	MovieID uint    `json:"movie_id"`
	UserID  string `json:"user_id"`
}
