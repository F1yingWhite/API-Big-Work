package models

import (
	"API_BIG_WORK/config"
	"log"

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
}

func Migrate() {
	DB.AutoMigrate(&Student{})
}
