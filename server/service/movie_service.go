package service

import (
	"API_BIG_WORK/models"
	"errors"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func DowFileWithOutAuth(c *gin.Context) {
	name := c.Param("name")
	_, err := models.GetMovieByPath(url.PathEscape("movies/" + name))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "文件不存在",
		})
		return
	}
	filename := path.Join("./movies", name)
	//响应一个文件
	c.File(filename)
}

type LikeMovie struct {
}

func (s *LikeMovie) Handle(c *gin.Context) (any, error) {
	id := c.GetString("id")
	//转为uint
	movie_id := c.Param("id")
	movieID, err := strconv.ParseUint(movie_id, 10, 32)
	if err != nil {
		return nil, err
	}
	err = models.LikeMovie(uint(movieID), id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type LikeMovieLeMe struct {
}

func (s *LikeMovieLeMe) Handle(c *gin.Context) (any, error) {
	id := c.GetString("id")
	//转为uint
	movie_id := c.Param("id")
	movieID, err := strconv.ParseUint(movie_id, 10, 32)
	if err != nil {
		return nil, err
	}
	_, err = models.GetLikeByIdAndMovieId(id, uint(movieID))
	if err != nil {
		return map[string]string{"like": "false"}, nil
	} else {
		return map[string]string{"like": "true"}, nil
	}
}

type GetMovieByAuthor struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func (s *GetMovieByAuthor) Handle(c *gin.Context) (any, error) {
	author, _ := c.Get("id")
	movies, err := models.GetMovieByAuthor(author.(string), s.Page, s.PageSize)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

type DeleteMovie struct {
}

func (s *DeleteMovie) Handle(c *gin.Context) (any, error) {
	movie_id := c.Param("id")
	movieID, err := strconv.ParseUint(movie_id, 10, 32)
	if err != nil {
		return nil, err
	}
	movie, err := models.GetMovieByID(int(movieID))
	if err != nil {
		return nil, err
	}
	if movie.Author != c.GetString("id") {
		return nil, errors.New("无权删除")
	}
	err = movie.Delete()
	if err != nil {
		return nil, err
	}
	//删除对应的文件	movie.Path
	os.Remove(path.Join("./", movie.Path))
	return nil, nil
}

type UploadMovie struct {
}

func (s *UploadMovie) Handle(c *gin.Context) (any, error) {
	videoName := c.PostForm("name")
	if videoName == "" {
		return nil, errors.New("视频名不能为空")
	}
	file, err := c.FormFile("video")
	if err != nil {
		return nil, err
	}
	//确认是不是mp4文件
	if path.Ext(file.Filename) != ".mp4" {
		return nil, errors.New("文件格式错误")
	}
	id := c.GetString("id")
	path := path.Join("movies", uuid.New().String()+".mp4")
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		return nil, err
	}
	return nil, models.CreateMovie(videoName, id, path)
}

type RecommendMovie struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func (s *RecommendMovie) Handle(c *gin.Context) (any, error) {
	id := c.GetString("id")
	RecommendMovies, err := models.RecommendMovie(id, s.Page, s.PageSize)
	if err != nil {
		return nil, err
	}
	return RecommendMovies, nil
}

type UpMovie struct {
	MovieId int `form:"movieId" binding:"required"`
}

func (s *UpMovie) Handle(c *gin.Context) (any, error) {
	// 找id比当前id小的最大的id
	movie, err := models.UpMovie(uint(s.MovieId))
	if err != nil {
		return nil, err
	}
	return movie, nil
}

type DownMovie struct {
	MovieId int `form:"movieId" binding:"required"`
}

func (s *DownMovie) Handle(c *gin.Context) (any, error) {
	// 找id比当前id大的最小的id
	movie, err := models.DownMovie(uint(s.MovieId))
	if err != nil {
		return nil, err
	}
	return movie, nil
}
