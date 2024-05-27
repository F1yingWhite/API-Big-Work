package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Like   int    `json:"like"`
	Path   string `json:"path"`
}

func CreateMovie(title string, author string, path string) error {
	movie := Movie{
		Title:  title,
		Author: author,
		Like:   0,
		Path:   path,
	}
	return DB.Create(&movie).Error
}

func CreateMovieWithLike(title string, author string, path string, like int) error {
	movie := Movie{
		Title:  title,
		Author: author,
		Like:   like,
		Path:   path,
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
