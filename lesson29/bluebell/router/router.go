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
	"GoWeb/lesson29/bluebell/logger"
	"GoWeb/lesson29/bluebell/middlewares"
	"GoWeb/lesson29/bluebell/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	// 设置 gin 框架日志输出模式
	gin.SetMode(settings.Conf.GinConfig.Mode)
	// 创建一个路由引擎
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务路由
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware()) // 应用 JWT 认证中间件
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.POST("/vote", controller.PostVoteHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
