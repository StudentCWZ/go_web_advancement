/*
   @Author: StudentCWZ
   @Description:
   @File: vote
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/24 09:49
*/

package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

/*
	投票的几种情况：
		direction=1 时，有两种情况：
			1. 之前没有投过票，现在投赞成票 --> 更新分数和投票记录 --> 差值的绝对值为 1
			2. 之前投反对票，现在改投赞成票 --> 更新分数和投票记录 --> 差值的绝对值为 2
		direction=0 时，有两种情况：
			1. 之前投反对票，现在要取消投票 --> 更新分数和投票记录 --> 差值的绝对值为 1
			2. 之前投赞成票，现在要取消投票 --> 更新分数和投票记录 --> 差值的绝对值为 1
		direction=-1 时，有两种情况：
			1. 之前没有投过票，现在投反对票 --> 更新分数和投票记录 --> 差值的绝对值为 1
			2. 之前投赞成票，现在改投反对票 --> 更新分数和投票记录 --> 差值的绝对值为 2

投票的限制：
	每个帖子自发表之日起一个星期内允许用户投票，超过一个星期就不允许再投票了。
		1. 到期之后将 redis 中保存的赞成票数及反对票数存储到 mysql 表中
		2. 到期之后删除那个 KeyPostScoreZSetPrefix
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票多少分
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
	ErrorVoteRepested   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) (err error) {
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 把帖子 id 加到社区的 set
	cKey := getRedisKey(KeyCommentZSetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err = pipeline.Exec()
	return
}

func VoteForPost(userID, postID string, value float64) (err error) {
	// 1. 判断投票的限制
	// 去 redis 取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	// 2. 更新帖子的分数
	// 2 和 3 需要放到一个 pipeline 事务中操作
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	// 如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrorVoteRepested
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3. 记录用户为该帖子投票的分数
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), postID)
		if err != nil {
			return
		}
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  0, // 赞成票还是反对票
			Member: nil,
		})
		if err != nil {
			return
		}
	}
	_, err = pipeline.Exec()
	return
}
