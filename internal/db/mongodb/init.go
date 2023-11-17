package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jjbot/core/config"
	"jjbot/core/logger"
)

var (
	Client = map[string]*api{}
)

type api struct {
	c *mongo.Client
}

func InitDB() {
	if len(config.C.Mongodb) == 0 {
		logger.SugarLogger.Fatalf("无法获得mongodb地址，终止程序")
		return
	}
	for _, c := range config.C.Mongodb {
		// 设置客户端连接配置
		clientOptions := options.Client().ApplyURI("mongodb://" + c.Addr)
		// 连接到MongoDB
		var err error
		cl, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			logger.SugarLogger.Fatal(err)
		}
		Client[c.Name] = &api{cl}
		// 检查连接
		err = Client[c.Name].c.Ping(context.TODO(), nil)
		if err != nil {
			Client[c.Name] = nil
			logger.SugarLogger.Fatal(err)
		}
		logger.SugarLogger.Infof("连接到MongoDB: %s", c.Name)
	}
}
