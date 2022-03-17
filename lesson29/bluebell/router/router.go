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
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
