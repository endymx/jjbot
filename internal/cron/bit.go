package cron

import (
	"jjbot/service/bot/bit"
	"jjbot/service/logger"
)

func BitCron() {
	crontab := NewCrontab()
	// 添加函数作为定时任务
	/*bitTaskFunc := func() {
		bit.UpdateFood()
	}
	if err := crontab.AddByFunc("bit", "0/10 * * * *", bitTaskFunc); err != nil {
		logger.SugarLogger.Fatalf("添加Task失败： %s", err)
	}*/

	bitTaskFunc2 := func() {
		bit.Command = []bool{false, false, false, false}
	}
	if err := crontab.AddByFunc("bit2", "0/1 * * * *", bitTaskFunc2); err != nil {
		logger.SugarLogger.Fatalf("添加Task失败： %s", err)
	}
	crontab.Start()
}
