/*
   @Author: StudentCWZ
   @Description:
   @File: redis
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 14:57
*/

package redis

import (
	"GoWeb/lesson25/web_app/settings"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// 声明一个全局的 rdb 变量
var rdb *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("Closed redis failed", zap.Error(err))
		return
	}
	return
}
