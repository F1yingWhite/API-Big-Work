package server

import (
	"API_BIG_WORK/server/middlewares"
	"API_BIG_WORK/server/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	api := r.Group("api")
	api.Use(gin.Recovery())
	api.Use(middlewares.RequestCounterMiddleware())
	api.Use(middlewares.RateLimitMiddleware(time.Second, 100, 10)) //一秒放10个令牌,就是一10qps
	api.Use(middlewares.Logger())
	{
		user := api.Group("user")
		{
			// POST api/user/login | 登录
			user.POST("login", service.HandlerBind(&service.Login{}))
			// POST api/user/register | 注册
			user.POST("register", service.HandlerBind(&service.Register{}))

			atuh := user.Group("")
			atuh.Use(middlewares.TokenAuthorization())
			{
				// GET api/user | 获取用户信息
				atuh.GET("", service.HandlerNoBind(&service.GetUser{}))
				// PUT api/user/password | 修改密码
				atuh.PUT("password", service.HandlerBind(&service.UpdatePassword{}))
				// PUT api/user/username | 修改用户名
				atuh.PUT("username", service.HandlerBind(&service.UpdateUserName{}))
			}
		}

	}
	return r
}
