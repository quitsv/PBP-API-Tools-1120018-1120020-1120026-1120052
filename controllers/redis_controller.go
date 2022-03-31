package controllers

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

var context_redis = context.Background()

func SetRedis(rdb *redis.Client, key string, value string, expiration int) {
	err := rdb.Set(context_redis, key, value, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetRedis(rdb *redis.Client, key string) string {
	val, err := rdb.Get(context_redis, key).Result()

	if err != nil {
		log.Fatal(err)
	}
	return val
}
