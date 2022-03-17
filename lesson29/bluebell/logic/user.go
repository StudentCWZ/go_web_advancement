/*
   @Author: StudentCWZ
   @Description:
   @File: user
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 16:57
*/

package logic

import (
	"GoWeb/lesson29/bluebell/dao/mysql"
	"GoWeb/lesson29/bluebell/pkg/snowflake"
)

// SignUp 存放业务逻辑代码
func SignUp() {
	// 1. 判断用户存不存在
	mysql.QueryUserByUsername()
	// 2. 生成 UID
	snowflake.GenID()
	// 3. 密码加密
	// 4. 保存进数据库
	mysql.InsertUser()
	// redis.xxx ...
}
