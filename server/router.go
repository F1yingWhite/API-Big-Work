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
			// PUT api/user/password | 修改密码
			user.PUT("password", service.HandlerBind(&service.UpdatePassword{}))

		}
		//需要jwt验证的接口
		auth := api.Group("")
		auth.Use(middlewares.TokenAuthorization())
		{
			users := auth.Group("user")
			{
				// GET api/user | 获取用户信息
				users.GET("", service.HandlerNoBind(&service.GetUser{}))
				// PUT api/user/username | 修改用户名
				users.PUT("username", service.HandlerBind(&service.UpdateUserName{}))
				// DELETE api/user | 删除用户
				users.DELETE("", service.HandlerNoBind(&service.DeleteUser{}))
			}

			movies := auth.Group("movie")
			{
				// GET api/movie | 获取指定id的电影
				movies.GET("", service.HandlerBindQuery(&service.GetMovie{}))
				// GET api/movie/list | 获取电影列表
				movies.GET("list", service.HandlerBindQuery(&service.GetMovieList{}))
				// GET api/movie/moviepath | 看电影
				movies.GET("/movies/:name", service.DowFile)
			}
		}
		// 静态文件服务，提供movie文件夹下的文件
	}
	return r
}
