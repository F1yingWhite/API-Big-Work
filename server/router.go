package server

import (
	"api2/server/middlewares"
	"api2/server/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MetricHttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "promdemo",
			Subsystem: "demo",
			Name:      "http_request_total",
			Help:      "http request total",
		},
		[]string{"from"},
	)
)

func InitRouter() *gin.Engine {
	r := gin.New()

	//一些基础配置
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := r.Group("api")
	api.Use(gin.Recovery())
	api.Use(middlewares.RequestCounterMiddleware())
	api.Use(middlewares.RateLimitMiddleware(time.Second, 100, 10)) //一秒放10个令牌,就是一10qps
	api.Use(middlewares.Logger())
	{
		//学生信息查询接口
		student := api.Group("student")
		{
			// GET /api/student?name=?&page=?&page_size=?&birth_start=?&birth_end=? | 查询学生信息
			student.GET("", service.HandlerBindQuery(&service.StudentService{}))
		}
	}
	return r
}
