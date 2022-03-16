/*
   @Author: StudentCWZ
   @Description:
   @File: redis
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 10:40
*/

package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 声明一个全局的 rdb 变量
var rdb *redis.Client

// Init 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
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
