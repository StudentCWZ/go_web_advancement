/*
   @Author: StudentCWZ
   @Description:
   @File: doc_response_models
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/25 10:20
*/

package controller

import "GoWeb/lesson29/bluebell/models"

// 专门用来存放接口文档用到的 model
// 因为我们的接口文档返回的数据格式是一致的，但是具体的 data 类型不一致

type _ResponsePostList struct {
	*_Response
	Data []*models.ApiPostDetail `json:"data"` // 数据
}

type _Response struct {
	Code    ResCode `json:"code"`    // 业务响应状态吗
	Message string  `json:"message"` // 提示信息
}

type _ResponseCommunityDetailList struct {
	*_Response
	Data []*models.CommunityDetail `json:"data"` // 数据
}

type _ResponseCommunityList struct {
	*_Response
	Data []*models.Community `json:"data"` // 数据
}
