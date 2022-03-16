/*
   @Author: StudentCWZ
   @Description:
   @File: routes
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 10:53
*/

package routes

import (
	"GoWeb/lesson24/web_app/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
