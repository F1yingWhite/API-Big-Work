package models

import (
	"api2/config"
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func postgresDB(dsn string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(1)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(cfg *config.Config) {
	var db *gorm.DB
	var err error

	db, err = postgresDB(cfg.Dsn, &gorm.Config{})

	if err != nil {
		log.Panicf("无法连接数据库，%s", err)
	}

	DB = db
	Migrate()
	//查询21301172学号的学生信息，如果为空，则创建10万条学生信息
	var student Student
	DB.Where("student_no = ?", 21301172).First(&student)
	if student.Id == 0 {
		CreateStudent()
	}
}

func Migrate() {
	DB.AutoMigrate(&Student{})
}

func CreateStudent() {
	for i := 0; i < 100000; i++ {
		//随机生成10万条学生信息，要求学号唯一，姓名随机，性别男/女，出生日期在2003~2010年之间
		StudentNo := 21301172 + i
		Name := "许" + fmt.Sprint(rand.Intn(100))
		var Gender string
		if rand.Intn(2) == 0 {
			Gender = "男"
		} else {
			Gender = "女"
		}
		Birth := time.Date(2003+rand.Intn(8), time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.Local)
		student := Student{
			StudentNo: StudentNo,
			Name:      Name,
			Gender:    Gender,
			Birth:     Birth,
		}
		DB.Create(&student)
	}
}
