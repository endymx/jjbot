package main

import (
	"jjbot/core/config"
	"jjbot/core/logger"
	"jjbot/core/v12"
)

func main() {
	logger.InitLogger() //初始化logger服务
	config.PullConf()   //读取conf配置文件，很重要勿注释
	//mongodb.InitDB()    //初始化数据库
	//redis.InitDB()
	v12.Create() //QQ机器人
	select {}
}
