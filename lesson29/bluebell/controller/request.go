/*
   @Author: StudentCWZ
   @Description:
   @File: request
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/21 16:59
*/

package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxtUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户 ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxtUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
