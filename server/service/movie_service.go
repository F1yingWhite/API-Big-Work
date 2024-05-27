package service

import (
	"API_BIG_WORK/models"

	"github.com/gin-gonic/gin"
)

type GetMovie struct {
	ID int `form:"id" binding:"required"`
}

func (s *GetMovie) Handle(c *gin.Context) (any, error) {
	movie, err := models.GetMovieByID(s.ID)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

type GetMovieList struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func (s *GetMovieList) Handle(c *gin.Context) (any, error) {
	movies, err := models.GetMovieList(s.Page, s.PageSize)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
