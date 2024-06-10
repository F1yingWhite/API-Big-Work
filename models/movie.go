package models

import (
	"net/url"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Like   uint   `json:"like" gorm:"default:0"`
	Path   string `json:"path"`
}

func CreateMovie(title string, author string, path string) error {
	movie := Movie{
		Title:  title,
		Author: author,
		Like:   0,
		Path:   url.PathEscape(path),
	}
	return DB.Create(&movie).Error
}

func CreateMovieWithLike(title string, author string, path string, like uint) error {

	movie := Movie{
		Title:  title,
		Author: author,
		Like:   like,
		Path:   url.PathEscape(path),
	}
	return DB.Create(&movie).Error
}

func GetMovieByID(id int) (Movie, error) {
	var movie Movie
	err := DB.Where("id = ?", id).First(&movie).Error
	return movie, err
}

func GetMovieByAuthor(author string, page, pageSize int) ([]Movie, error) {
	var movies []Movie
	err := DB.Where("author = ?", author).Offset((page - 1) * pageSize).Limit(pageSize).Find(&movies).Error
	return movies, err
}

func GetMovieByTitle(title string, page, pageSize int) ([]Movie, error) {
	var movies []Movie
	err := DB.Where("title = ?", title).Offset((page - 1) * pageSize).Limit(pageSize).Find(&movies).Error
	return movies, err
}

func GetMovieList(page, pageSize int) ([]Movie, error) {
	var movies []Movie
	err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&movies).Error
	return movies, err
}

func GetMovieByPath(path string) (Movie, error) {
	var movie Movie
	err := DB.Where("path = ?", path).First(&movie).Error
	return movie, err
}

// 推荐电影 每次推荐一个，推荐like最高的，如果like相同，推荐最新的，用户已经看过的不推荐
func RecommendMovie(userID string) (Movie, error) {
	var movie Movie
	err := DB.Raw(`
	SELECT movies.* FROM movies
	LEFT JOIN histories ON movies.id = histories.movie_id AND histories.user_id = ?
	WHERE histories.movie_id IS NULL
	ORDER BY movies.like DESC, movies.created_at DESC
	LIMIT 1`, userID).Scan(&movie).Error
	return movie, err
}

func DeleteMovie(id uint) error {
	return DB.Where("id = ?", id).Delete(&Movie{}).Error
}

func (m *Movie) Delete() error {
	return DB.Delete(m).Error
}
