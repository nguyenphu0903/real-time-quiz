package pkg

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
}

func GetRedisCtx() context.Context {
	return context.Background()
}
