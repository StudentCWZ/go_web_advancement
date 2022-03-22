/*
   @Author: StudentCWZ
   @Description:
   @File: error_code
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/22 16:31
*/

package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在, 请勿重复注册！")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的 ID")
)
