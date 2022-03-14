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

// watch 操作
func watchRedisExample() {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 业务逻辑
			time.Sleep(time.Second * 5) // 演示其他客户端修改 key 的情况，导致事务失败
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("tx exec failed, err: %v\n", err)
		return
	}
	fmt.Println("tx exec success")
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
	// watch 操作
	watchRedisExample()
}
