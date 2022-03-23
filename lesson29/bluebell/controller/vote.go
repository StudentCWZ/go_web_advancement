/*
   @Author: StudentCWZ
   @Description:
   @File: vote
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/23 17:01
*/

package controller

import (
	"GoWeb/lesson29/bluebell/logic"
	"GoWeb/lesson29/bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 投票

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamsVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	logic.PostVote()
	ResponseSuccess(c, nil)
}
