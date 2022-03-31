package controllers

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

var context_redis = context.Background()

func rdbSet() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	SetRedis(rdb, "diskon", "Beli sekarang diskon 50%", 0)
}

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
