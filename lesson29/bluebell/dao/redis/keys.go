/*
   @Author: StudentCWZ
   @Description:
   @File: keys
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/23 16:36
*/

package redis

// redis key
// redis key 注意使用命名空间的方式，方便查询和拆分

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // ZSet: 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // ZSet: 帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" // ZSet: 记录用户及投票类型，参数是 post_id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
