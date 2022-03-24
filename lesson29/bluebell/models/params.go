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
const (
	OrderTime  = "Time"
	OrderScore = "score"
)

// ParamsSignUp 注册请求参数
type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamsLogin 登录请求参数
type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamsVoteData 投票数据
type ParamsVoteData struct {
	//UserID 从请求中获取当前的用户
	PostID    int64 `json:"post_id,string" binding:"required"`       // 帖子 id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票/反对票/取消投票
}

// ParamsPostList 获取帖子列表 query string 参数
type ParamsPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
