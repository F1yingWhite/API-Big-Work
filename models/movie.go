package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
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

func GetMovieByID(id int) (Movie, error) {
	var movie Movie
	err := DB.Where("id = ?", id).First(&movie).Error
	return movie, err
}

func GetMovieByAuthor(author string) ([]Movie, error) {
	var movies []Movie
	err := DB.Where("author = ?", author).Find(&movies).Error
	return movies, err
}
