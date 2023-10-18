package redis

import (
	"github.com/redis/go-redis/v9"
	"jjbot/service/conf"
	"jjbot/service/logger"
)

var Client *redis.Client

func InitDB() {
	Client = redis.NewClient(&redis.Options{
		Addr:     conf.C.Redis.Addr,
		Password: conf.C.Redis.Password,
		DB:       0,
	})
	logger.SugarLogger.Info("尝试连接到Redis...")
}
