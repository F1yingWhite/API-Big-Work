package config

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Dsn     string `yaml:"dsn"`
	LogFile string `yaml:"log_file"`
}

func ReadConfig() (*Config, error) {
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		f, err := os.Create("config.yaml")
		if err != nil {
			return nil, err
		}

		config := &Config{
			Dsn:     "host=localhost user=postgres password=123 dbname=API2 port=5432 sslmode=disable",
			LogFile: "log.txt",
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

	return &cfg, nil
}

func InitLog(config *Config) {
	f, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = f
}
