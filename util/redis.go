package util

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_URI,
		Password: "", // no password set
		DB:       0,  // use default MongoClient
	})
}

func RedisSet(key string, value interface{}, ttl int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := RedisClient.Set(ctx, key, value, time.Duration(ttl)*time.Minute).Err()
	if err != nil {
		panic(err)
	}
}

func RedisGet(key string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	return val
}
