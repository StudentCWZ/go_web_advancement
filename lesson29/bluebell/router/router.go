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
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
