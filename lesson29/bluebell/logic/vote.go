/*
   @Author: StudentCWZ
   @Description:
   @File: vote
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/23 16:54
*/

package logic

import (
	"GoWeb/lesson29/bluebell/dao/redis"
	"GoWeb/lesson29/bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能
// 本项目使用简化版的投票分数
// 投一票就加 432 分 86400/200 —> 需要 200 张赞成票可以给你的帖子续一天

// 1. 判断投票的限制
// 2. 更新帖子的分数
// 3. 记录用户为该帖子投票的分数

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamsVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.Int64("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))

}
