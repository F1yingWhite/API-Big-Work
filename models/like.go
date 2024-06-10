package models

// 记录用户点赞记录，已经点赞过的不能再点赞
type Like struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	MovieID uint   `json:"movie_id"`
	UserID  string `json:"user_id"`
}

func LikeMovie(id uint, userID string) error {
	//先查询有没有点赞记录,如果没有就创建,有就删除这次点赞
	var like Like
	err := DB.Where("movie_id = ? AND user_id = ?", id, userID).First(&like).Error
	if err != nil {
		like = Like{
			MovieID: id,
			UserID:  userID,
		}
		return DB.Create(&like).Error
	}
	return DB.Delete(&like).Error
}
