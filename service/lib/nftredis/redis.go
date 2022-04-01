package nftredis

import (
	"github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "capregsoft",
	})
	return client
}
