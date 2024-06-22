package service

import (
	"API_BIG_WORK/config"
	"API_BIG_WORK/models"
	"API_BIG_WORK/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type Login struct {
	ID       string `form:"id"`
	Password string `form:"password"`
}

func (s *Login) Handle(c *gin.Context) (any, error) {
	user, err := models.GetUserByID(s.ID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if !utils.ComparePasswords(user.Password, s.Password) {
		return nil, errors.New("密码错误")
	}
	//得到token
	token, err := utils.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}
	// 将token存入redis
	if err := models.Redis.Set(token, user.ID, time.Duration(config.CFG.JWTExpire)).Err(); err != nil {
		return nil, err
	}
	return []map[string]interface{}{
		{
			"token": token,
		},
	}, nil
}

type Register struct {
	ID       string `form:"id"`
	Name     string `form:"name"`
	Password string `form:"password"`
}

func (s *Register) Handle(c *gin.Context) (any, error) {
	//密码复杂度验证,必须包含大小写字母和数字且长度大于等于8
	if len(s.Password) < 8 {
		return nil, errors.New("密码长度必须大于等于8")
	}
	if !utils.VerifyPassword(s.Password) {
		return nil, errors.New("密码必须包含大小写字母和数字")
	}
	if err := models.CreateUser(s.ID, s.Name, s.Password); err != nil {
		return nil, errors.New("用户已存在")
	}
	return nil, nil
}

type GetUser struct {
}

func (s *GetUser) Handle(c *gin.Context) (any, error) {
	id, _ := c.Get("id")
	user, err := models.GetUserByID(id.(string))
	if err != nil {
		return nil, err
	}
	return user, nil
}

type UpdatePassword struct {
	Id          string `form:"id"`
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}

func (s *UpdatePassword) Handle(c *gin.Context) (any, error) {
	err := models.UpdateUserPassword(s.Id, s.OldPassword, s.NewPassword)
	return nil, err
}

type UpdateUserName struct {
	Name string `form:"name"`
}

func (s *UpdateUserName) Handle(c *gin.Context) (any, error) {
	id, _ := c.Get("id")
	err := models.UpdateUserUsername(id.(string), s.Name)
	return nil, err
}

type DeleteUser struct{}

func (s *DeleteUser) Handle(c *gin.Context) (any, error) {
	id, _ := c.Get("id")
	err := models.DeleteUser(id.(string))
	return nil, err
}
