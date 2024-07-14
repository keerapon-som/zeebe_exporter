package utils

import (
	"github.com/go-redis/redis/v8"
)

// Assuming you have a Redis client set up as follows:

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",       // Redis server address
		Password: "exampleRedisPassword", // no password set
		DB:       0,                      // use default DB
	})
}

func GetRDB() *redis.Client {
	if rdb == nil {
		InitRedis()
	}
	return rdb
}
