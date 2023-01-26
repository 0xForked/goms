package main

import (
	"context"
	"fmt"
	"github.com/aasumitro/goms/internal/notify"
	"github.com/go-redis/redis/v8"
	"sync"
)

const RedisAddress = "localhost:6379"

var (
	redisOnce sync.Once
	redisPool *redis.Client
)

func init() {
	getRedisConn()
}

func main() {
	notify.NewNotifyService(redisPool)
}

func getRedisConn() {
	redisOnce.Do(func() {
		redisPool = redis.NewClient(&redis.Options{Addr: RedisAddress})
		if err := redisPool.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf(
				"REDIS_ERROR: %s",
				err.Error()))
		}
	})
}
