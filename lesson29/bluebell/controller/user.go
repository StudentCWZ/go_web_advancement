/*
   @Author: StudentCWZ
   @Description:
   @File: user
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 16:51
*/

package controller

import (
	"GoWeb/lesson29/bluebell/dao/mysql"
	"GoWeb/lesson29/bluebell/logic"
	"GoWeb/lesson29/bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
// @Summary 用户注册请求接口
// @Description 用户注册请求接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamsSignUp false "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} _Response
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamsSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断 error 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidPassword, removeTopStruct(errs.Translate(trans))) // 翻译错误
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	// 请求参数有误，直接返回响应
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Printf("user: %#v\n", *p)
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		// 记录错误日志
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理用户登录的函数
// @Summary 用户登录请求接口
// @Description 用户登录请求接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body models.ParamsLogin false "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数和参数校验
	p := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断 error 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) // 翻译错误
		return
	}
	// 2. 业务处理
	user, err := logic.Login(p)
	if err != nil {
		// 记录错误日志
		zap.L().Error("Logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		//
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID), // 前端 id 值最大为 2^53-1；后端 int64 最大值为 2^63-1；不一致可能会导致失真
		"username": user.Username,
		"token":    user.Token,
	})
}
