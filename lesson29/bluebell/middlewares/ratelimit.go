/*
@Time : 2022/3/26 21:39
@Author : StudentCWZ
@File : ratelimit
@Software: GoLand
*/

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit ....")
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
