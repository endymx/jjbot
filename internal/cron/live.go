package cron

import (
	"jjbot/service/bot/live"
	"jjbot/service/logger"
)

func LiveCron() {
	crontab := NewCrontab()
	/*
		// 实现接口的方式添加定时任务
		LiveTask := &task.LiveTask{}
		if err := crontab.AddByID("1", "0/1 * * * *", LiveTask); err != nil {
			logger.SugarLogger.Fatalf("添加Task失败： %s", err)
		}
	*/

	// 添加函数作为定时任务
	liveTaskFunc := func() {
		live.GetLiveInfo()
	}
	if err := crontab.AddByFunc("live", "0/1 * * * *", liveTaskFunc); err != nil {
		logger.SugarLogger.Fatalf("添加Task失败： %s", err)
	}

	crontab.Start()
}
