package model

import (
	"antiNCP/config"
	. "antiNCP/log"
	"github.com/go-redis/redis/v7"
)

// 使用文档 https://redis.uptrace.dev/#executing-commands

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     config.C.Redis.Addr,
			Password: config.C.Redis.Password,
			DB:       config.C.Redis.Db,
		},
	)

	err := redisClient.Ping().Err()
	if err != nil {
		Logger.Panic(err)
	}
	Logger.Println("Successfully connected to Redis.")
}
