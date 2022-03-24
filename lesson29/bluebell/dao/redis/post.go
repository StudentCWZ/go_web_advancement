/*
   @Author: StudentCWZ
   @Description:
   @File: post
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/24 14:22
*/

package redis

import (
	"GoWeb/lesson29/bluebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	// 确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1
	// ZRevRange 查询：按照分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamsPostList) ([]string, error) {
	// 从 redis 获取 id
	// 根据用户请求中携带的 order 参数确定要查询的 redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFormKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据 ids 查询每篇帖子投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	// 查找 key 中分数是 1 的元素的数量 --> 统计每篇帖子的赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	// 使用 pipeline 一次发送多条命令，减少 RTT
	pipeline := client.TxPipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmderData, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmderData {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 按社区查询 ids
func GetCommunityPostIDsInOrder(p *models.ParamsCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用 ZInterStore 把分区的帖子 set 与帖子分数的 zSet 生成一个新的 zSet
	// 针对的新的 ZSet 按之前的逻辑取数
	// 利用缓存 key 减少 ZInterStore 执行的次数
	// 社区的 key
	cKey := getRedisKey(KeyCommentZSetPrefix + strconv.Itoa(int(p.CommunityID)))
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(orderKey).Val() < 1 {
		// 不存在，需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "Max",
		}, cKey, orderKey) // ZInterStore 计算
		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在的话就直接根据 key 查询 ids
	return getIDsFormKey(key, p.Page, p.Size)
}
