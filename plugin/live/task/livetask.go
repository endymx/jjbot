package task

import (
	"jjbot/service/bot/live"
)

type LiveTask struct {
}

func (t *LiveTask) Run() {
	live.GetLiveInfo()
}
