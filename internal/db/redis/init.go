package redis

import (
	"github.com/redis/go-redis/v9"
	"jjbot/core/config"
	"jjbot/core/logger"
)

var Client *redis.Client

func InitDB() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Addr,
		Password: config.C.Redis.Password,
		DB:       0,
	})
	logger.SugarLogger.Info("尝试连接到Redis...")
}
