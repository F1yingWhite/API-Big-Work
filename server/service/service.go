package service

import (
	"API_BIG_WORK/utils"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	None = 0
	Bind = 1 << iota
	BindUri
	BindQuery
)

type Service interface {
	Handle(c *gin.Context) (any, error)
}

func HandlerNoBind(s Service) gin.HandlerFunc {
	return HandlerWithBindType(s, None)
}

func HandlerBind(s Service) gin.HandlerFunc {
	return HandlerWithBindType(s, Bind)
}

func HandlerBindUri(s Service) gin.HandlerFunc {
	return HandlerWithBindType(s, BindUri)
}

func HandlerBindQuery(s Service) gin.HandlerFunc {
	return HandlerWithBindType(s, BindQuery)
}

func HandlerWithBindType(s Service, bindType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		LogID := c.GetString("logID")
		if bindType&BindUri != 0 {
			if err = c.ShouldBindUri(s); err != nil {
				response := utils.ErrorResponse(errors.New("参数类型错误"))
				log.Printf("[LogID:%s] Response: %s", LogID, response.ErrorStr)
				c.JSON(http.StatusBadRequest, response)
				return
			}
		}
		if bindType&Bind != 0 {
			if err = c.ShouldBind(s); err != nil {
				response := utils.ErrorResponse(errors.New("参数类型错误"))
				log.Printf("[LogID:%s] Response: %s", LogID, response.ErrorStr)
				c.JSON(http.StatusBadRequest, response)
				return
			}
		}
		if bindType&BindQuery != 0 {
			if err = c.ShouldBindQuery(s); err != nil {
				response := utils.ErrorResponse(errors.New("参数类型错误"))
				log.Printf("[LogID:%s] Response: %s", LogID, response.ErrorStr)
				c.JSON(http.StatusBadRequest, response)
				return
			}
		}

		res, err := s.Handle(c)
		if err != nil {
			response := utils.ErrorResponse(errors.New("参数类型错误"))
			log.Printf("[LogID:%s] Response: %s", LogID, response.ErrorStr)
			c.JSON(http.StatusBadRequest, response)
		} else {
			response := utils.Response(res)
			log.Printf("[LogID:%s] Response: %v", LogID, response)
			c.JSON(http.StatusOK, response)
		}
	}
}
