/*
   @Author : cuiweizhi
*/

package main

import (
	"fmt"
	"github.com/go-redis/redis"
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

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err: %v\n", err)
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err: %v\n", err)
		return
	}
	fmt.Println("score:", val)

	varTwo, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err: %v\n", err)
	} else {
		fmt.Println("name:", varTwo)
	}
}

func hgetRedisExample() {
	val, err := rdb.HGetAll("user").Result()
	if err == redis.Nil {
		fmt.Println("user does not exist")
	} else if err != nil {
		fmt.Printf("get user failed, err: %v\n", err)
		return
	} else {
		fmt.Println("user:", val)
	}
	valTwo, err := rdb.HMGet("user", "name", "age").Result()
	if err == redis.Nil {
		fmt.Println("user does not exist")
	} else if err != nil {
		fmt.Printf("hmget user failed, err: %v\n", err)
		return
	} else {
		fmt.Println(valTwo)
	}
	valThree, err := rdb.HGet("user", "age").Result()
	if err == redis.Nil {
		fmt.Println("the field of age dose not exit or user does not exist")
	} else if err != nil {
		fmt.Printf("hget user.age failed, err: %v\n", err)
		return
	} else {
		fmt.Println("user.age:", valThree)
	}
}

func zsetRedisExample() {
	zsetkey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetkey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err: %v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把 Golang 的分数加 10
	newScore, err := rdb.ZIncrBy(zsetkey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err: %v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高 3 个
	ret, err := rdb.ZRevRangeWithScores(zsetkey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err: %v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取 95～100 分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetkey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err: %v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
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
	// 执行 redis 操作
	redisExample()
	// 执行 redis 操作
	hgetRedisExample()
	// 执行 redis 操作
	zsetRedisExample()
}
