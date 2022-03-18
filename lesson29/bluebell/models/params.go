/*
   @Author: StudentCWZ
   @Description:
   @File: params
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/18 08:32
*/

package models

// 定义请求的参数结构体

type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
