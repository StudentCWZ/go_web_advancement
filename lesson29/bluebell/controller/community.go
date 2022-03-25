/*
   @Author: StudentCWZ
   @Description:
   @File: community
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/22 15:10
*/

package controller

import (
	"GoWeb/lesson29/bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 社区详情
// @Summary 获取社区详情接口
// @Description 获取社区详情接口
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityList
// @Router /community [get]
// ---- 跟社区相关的 ----
func CommunityHandler(c *gin.Context) {
	// 查询到所有社区(community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
// @Summary 获取社区分类详情接口
// @Description 获取社区分类详情接口
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query string false "社区 id 参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityDetailList
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区 id
	idStr := c.Param("id") // 获取路径参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据 id 获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
