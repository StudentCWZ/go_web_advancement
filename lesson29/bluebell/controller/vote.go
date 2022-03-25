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
	"go.uber.org/zap"
)

// PostVoteHandler 投票
// @Summary 投票接口
// @Description 投票接口
// @Tags 投票相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body models.ParamsVoteData false "投票参数"
// @Security ApiKeyAuth
// @Success 200 {object} _Response
// @Router /vote [post]
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
	// 获取当前请求的用户 id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
