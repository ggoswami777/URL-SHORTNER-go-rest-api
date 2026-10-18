package config

import (
	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct{
	Client *redis.Client
}

func NewRedisClient() *RedisClient{
	rdb:=redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	

	return &RedisClient{
		Client:rdb,
	}
}