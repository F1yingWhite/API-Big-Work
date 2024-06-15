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

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
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

func sqliteDB(config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("movie.db"), config)
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
	if config.Postgresql {
		db, err = postgresDB(cfg.Dsn, &gorm.Config{})
	} else {
		db, err = sqliteDB(&gorm.Config{})
	}

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
	DB.AutoMigrate(&Like{})
}

func moveFile(sourcePath, destPath string) error {
	err := os.Rename(sourcePath, destPath)
	if err != nil {
		return err
	}
	return nil
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
	movieDir := "movies"
	var mp4Files []string

	err = filepath.Walk(movieDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查文件扩展名是否为 .mp4
		if !info.IsDir() && filepath.Ext(path) == ".mp4" {
			mp4Files = append(mp4Files, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	for _, mp4File := range mp4Files {
		new_id := uuid.New().String()
		new_path := fmt.Sprintf("movies/%s.mp4", new_id)
		err := moveFile(mp4File, new_path)
		if err != nil {
			log.Println(err)
		}
		filenameWithExtension := filepath.Base(mp4File)
		filename := strings.TrimSuffix(filenameWithExtension, ".mp4")
		CreateMovieWithLike(filename, fmt.Sprintf("user%d", r.Intn(20)), new_path, uint(rand.Intn(200000)))
	}
}
