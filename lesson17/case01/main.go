/*
   @Author : cuiweizhi
*/

package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 声明一个全局的 rdb 变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379", // docker 的 redis
		Password: "",                // no password set
		DB:       0,                 // use default DB
		PoolSize: 100,               // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return err
}

// 在某些场景下，当我们有多条命令要执行时，就可以考虑使用 pipeline 来优化
func pipelineRedisExample() {
	pipe := rdb.Pipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)
	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed, err: %v\n", err)
		return
	}
	fmt.Println("connect redis success")
	// 程序退出时释放相关资源
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			fmt.Printf("close redis failed, err: %v\n", err)
			return
		}
	}(rdb)
	// pipeline
	pipelineRedisExample()
}
