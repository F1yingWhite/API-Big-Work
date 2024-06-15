package main

import (
	"API_BIG_WORK/config"
	"API_BIG_WORK/models"
	"API_BIG_WORK/server"
	"API_BIG_WORK/server/middlewares"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)

func runPnpmDev() {
	// 执行 PNPM dev 命令
	cmd := exec.Command("pnpm", "--dir", "tiktok-web", "dev")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {
	config.InitFlag()
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Panicln(err)
	}
	config.InitLog(cfg)
	models.InitDB(cfg)
	if config.Dev {
		go runPnpmDev()
	}
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
