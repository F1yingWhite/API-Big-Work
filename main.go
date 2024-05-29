package main

import (
	"API_BIG_WORK/config"
	"API_BIG_WORK/models"
	"API_BIG_WORK/server"
	"API_BIG_WORK/server/middlewares"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

func Init() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Panicln(err)
	}
	config.InitLog(cfg)
	models.InitDB(cfg)

	go func() {
		for {
			time.Sleep(time.Second)
			requestsLastSecond := atomic.SwapInt64(&middlewares.RequestCounter, 0)
			fmt.Printf("\rRequests in last second: %d", requestsLastSecond)
		}
	}()
}

func main() {
	Init()
	api := server.InitRouter()

	err := api.Run(":8888")
	if err != nil {
		log.Panicln(err)
	}
}

//git连接测试
