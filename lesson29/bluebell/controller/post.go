/*
   @Author: StudentCWZ
   @Description:
   @File: post
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/23 08:56
*/

package controller

import (
	"GoWeb/lesson29/bluebell/logic"
	"GoWeb/lesson29/bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子
// @Summary 创建帖子接口
// @Description 创建帖子接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body models.Post false "创建帖子参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /post [post]
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(&p) error", zap.Any("err", err))
		zap.L().Error("create post with in invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 中取到当前发请求的用户的 ID 值
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
// @Summary 获取帖子详情接口
// @Description 获取帖子详情接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query string false "帖子 id 参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post/:id [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数以及参数校验(帖子的 id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 根据 id 取出帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 基础版获取帖子列表的处理函数
// @Summary 基础版帖子列表接口
// @Description 分页查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param page query string false "查询参数"
// @Param size query string false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts [get]
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 1. 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		return
	}
	// 2. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListTwoHandler 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamsPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListTwoHandler(c *gin.Context) {
	// GET 请求参数：/api/v1/post2?page=1&size=10&order=time
	// 获取分页参数
	// 初始化结构体时指定初始参数
	p := &models.ParamsPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	//c.ShouldBindJSON() // 如果请求中携带的是 json 格式数据，才能用这个方法获取到数据
	// c.ShouldBind() // 根据请求的数据类型选择相应的方法去获取数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListTwoHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 1. 获取数据
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2. 返回响应
	ResponseSuccess(c, data)
}

// GetCommunityPostListHandler 根据社区去查询帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	// 初始化结构体时指定初始参数
//	p := &models.ParamsCommunityPostList{
//		ParamsPostList: &models.ParamsPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	//c.ShouldBindJSON() // 如果请求中携带的是 json 格式数据，才能用这个方法获取到数据
//	// c.ShouldBind() // 根据请求的数据类型选择相应的方法去获取数据
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	// 1. 获取数据
//	data, err := logic.GetCommunityPostListHandler(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		return
//	}
//	// 2. 返回响应
//	ResponseSuccess(c, data)
//}
