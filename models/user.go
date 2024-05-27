package models

import (
	"API_BIG_WORK/config"
	"API_BIG_WORK/utils"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
}

func (u *User) Delete() error {
	return DB.Delete(u).Error
}

func DeleteUser(ID string) error {
	return DB.Where("id = ?", ID).Delete(&User{}).Error
}

func CreateUser(ID string, username string, password string) error {
	password, err := utils.HashPassword(password, config.CFG.Salt)
	if err != nil {
		return err
	}
	user := User{
		ID:       ID,
		Username: username,
		Password: password,
	}
	return DB.Create(&user).Error
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserByID(id string) (User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func UpdateUserPassword(ID string, old_password string, new_password string) error {
	if user, err := GetUserByID(ID); err != nil {
		return err
	} else if !utils.ComparePasswords(user.Password, old_password, config.CFG.Salt) {
		return errors.New("密码错误")
	} else {
		if password, err := utils.HashPassword(new_password, config.CFG.Salt); err != nil {
			return err
		} else {
			return DB.Model(&user).Update("password", password).Error
		}
	}
}

func UpdateUserUsername(ID string, new_username string) error {
	if _, err := GetUserByUsername(new_username); err == nil {
		return errors.New("用户名已存在")
	}
	if user, err := GetUserByID(ID); err != nil {
		return err
	} else {
		return DB.Model(&user).Update("username", new_username).Error
	}
}
