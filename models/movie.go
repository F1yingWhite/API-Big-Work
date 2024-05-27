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

func LikeMovie(id uint) error {
	return DB.Model(&Movie{}).Where("id = ?", id).Update("like", gorm.Expr("like + ?", 1)).Error
}

func DeleteMovie(id uint) error {
	return DB.Where("id = ?", id).Delete(&Movie{}).Error
}

func (m *Movie) Delete() error {
	return DB.Delete(m).Error
}
