package main

import (
	cron2 "jjbot/common/cron"
	"jjbot/common/db/mongodb"
	"jjbot/common/db/redis"
	"jjbot/service/bot"
	"jjbot/service/conf"
	"jjbot/service/logger"
)

func main() {
	logger.InitLogger() //初始化logger服务
	conf.PullConf()     //读取conf配置文件，很重要勿注释
	mongodb.InitDB()    //初始化数据库
	redis.InitDB()
	bot.Create()     //QQ机器人
	cron2.LiveCron() //启动cron定时任务
	cron2.BitCron()
}
