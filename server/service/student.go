package service

import (
	"API_BIG_WORK/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// 1. (20分) 使用实现学生信息查询接口(http接口)，支持按照姓名前缀或者birth范围查询，pageSize为1~1000的int值，返回相应数量的学生信息
type StudentService struct {
	Page       int    `json:"page" form:"page" binding:"required"`
	PageSize   int    `json:"page_size" form:"page_size" binding:"required"`
	BirthStart string `json:"birth_start" form:"birth_start"`
	BirthEnd   string `json:"birth_end" form:"birth_end"`
}

func (s *StudentService) Handle(c *gin.Context) (any, error) {
	if s.PageSize < 1 {
		return nil, errors.New("page_size should be bigger than 0")
	}
	if s.PageSize > 1000 {
		s.PageSize = 1000
	}
	if s.Page < 1 {
		return nil, errors.New("page should be bigger than 0")
	}
	LogID := c.GetString("logID")
	name := c.Query("name")
	birth_start := c.Query("birth_start")
	birth_end := c.Query("birth_end")
	if birth_start != "" && birth_end != "" {
		start, err := time.Parse("2006-01-02", s.BirthStart)
		if err != nil {
			return nil, errors.New("birth_start format error, should be like 2006-01-02")
		}
		end, err := time.Parse("2006-01-02", s.BirthEnd)
		if err != nil {
			return nil, errors.New("birth_end format error, should be like 2006-01-02")
		}
		if name != "" {
			return models.GetStudentByBirthRangeAndName(name, start, end, s.Page, s.PageSize, LogID)
		} else {
			return models.GetStudentByBirth(start, end, s.Page, s.PageSize, LogID)
		}
	}
	if birth_start == "" && birth_end != "" {
		end, err := time.Parse("2006-01-02", s.BirthEnd)
		if err != nil {
			return nil, errors.New("birth_end format error, should be like 2006-01-02")
		}
		if name != "" {
			return models.GetStudentByNameAndBirthLessThan(name, end, s.Page, s.PageSize, LogID)
		} else {
			return models.GetStudentByBirthLessThan(end, s.Page, s.PageSize, LogID)
		}
	}
	if birth_start != "" && birth_end == "" {
		start, err := time.Parse("2006-01-02", s.BirthStart)
		if err != nil {
			return nil, errors.New("birth_start format error, should be like 2006-01-02")
		}
		if name != "" {
			return models.GetStudentByNameAndBirthBiggerThan(name, start, s.Page, s.PageSize, LogID)
		} else {
			return models.GetStudentByBirthBiggerThan(start, s.Page, s.PageSize, LogID)
		}
	}
	if name != "" {
		return models.GetStudentByName(name, s.Page, s.PageSize, LogID)
	}
	return models.GetStudent(s.Page, s.PageSize, LogID)
}
