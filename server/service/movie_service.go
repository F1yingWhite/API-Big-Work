package service

import (
	"API_BIG_WORK/models"
	"net/url"
	"path"

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

func DowFile(c *gin.Context) {
	id, _ := c.Get("id")
	name := c.Param("name")
	//创建历史 movies%2F352b9997-c1de-49d0-a546-8ad69cd8253b.mp4
	movie, err := models.GetMovieByPath(url.PathEscape("movies/" + name))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "文件不存在",
		})
		return
	}
	models.CreateHistory(movie.ID, id.(string))
	filename := path.Join("./movies", name)
	//响应一个文件
	c.File(filename)
}
