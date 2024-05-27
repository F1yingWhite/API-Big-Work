package models

import (
	"API_BIG_WORK/config"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
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
	InitData()
}

func Migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Movie{})
	DB.AutoMigrate(&History{})
}

func InitData() {
	//查询用户0是否存在，如果存在则不再创建
	if _, err := GetUserByID("user0"); err == nil {
		return
	}
	//创建20个用户
	for i := 0; i < 20; i++ {
		CreateUser(fmt.Sprintf("user%d", i), fmt.Sprintf("user%d", i), "123456")
	}
	//遍历movie文件夹，读取里面的mp4视频，并且生成movie数据
	_, err := os.Stat("movies")
	if os.IsNotExist(err) {
		os.MkdirAll("movies", os.ModePerm)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	filepath.Walk("movies", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".mp4") {
			CreateMovieWithLike(strings.TrimSuffix(info.Name(), ".mp4"), fmt.Sprintf("user%d", r.Intn(21)), path, r.Intn(200000))
		}
		return nil
	})
}
