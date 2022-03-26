/*
   @Author: StudentCWZ
   @Description:
   @File: router
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 16:13
*/

package router

import (
	"GoWeb/lesson29/bluebell/controller"
	_ "GoWeb/lesson29/bluebell/docs" // 千万不要忘了导入上一步生成的docs
	"GoWeb/lesson29/bluebell/logger"
	"GoWeb/lesson29/bluebell/middlewares"
	"GoWeb/lesson29/bluebell/settings"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func Setup() *gin.Engine {
	// 设置 gin 框架日志输出模式
	gin.SetMode(settings.Conf.GinConfig.Mode)
	// 创建一个路由引擎
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 注册 swagger api 相关路由
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 注册业务路由
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Second, 1)) // 应用 JWT 认证中间件以及令牌桶限流中间件
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListTwoHandler)
		v1.POST("/vote", controller.PostVoteHandler)
	}
	pprof.Register(r) // 注册 pprof 相关路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
