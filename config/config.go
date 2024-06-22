package config

import (
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var CFG Config

type Config struct {
	PostgresqlDsn    string `yaml:"postgresql_dsn"`
	LogFile          string `yaml:"log_file"`
	JWTSigningString string `yaml:"jwt_signing_string"`
	RedisDsn         string `yaml:"redis_dsn"`
	RedisPassword    string `yaml:"redis_password"`
	RedisDB          int    `yaml:"redis_db"`
	JWTExpire        int    `yaml:"jwt_expire"`
}

func ReadConfig() (*Config, error) {
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		f, err := os.Create("config.yaml")
		if err != nil {
			return nil, err
		}

		config := &Config{
			PostgresqlDsn:    "host=localhost user=postgres password=1234 dbname=API_BIG_WORK port=5432 sslmode=disable",
			LogFile:          "log.txt",
			JWTSigningString: GenerateRandomString(32),
			RedisDsn:         "localhost:6379",
			RedisPassword:    "",
			RedisDB:          0,
			JWTExpire:        3600,
		}

		data, err := yaml.Marshal(&config)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(data)
		if err != nil {
			return nil, err
		}

		err = f.Close()
		if err != nil {
			return nil, err
		}
	}
	f, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}
	CFG = cfg
	return &cfg, nil
}

func InitLog(config *Config) {
	f, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	gin.DefaultWriter = f
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
