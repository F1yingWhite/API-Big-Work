package models

// 记录用户点赞记录，已经点赞过的不能再点赞
type Like struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	MovieID uint   `json:"movie_id"`
	UserID  string `json:"user_id"`
}
