package main

import (
	"api2/config"
	"api2/models"
	"api2/server"
	"api2/server/middlewares"
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
