package util

import (
	"context"

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

func RedisSet(key string, value string) {
	ctx := context.Background()
	err := RedisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}
