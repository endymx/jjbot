package ys

type Daily struct {
	Status
	Data DailyData `json:"data"`
}

type DailyData struct {
	CurrentResin              int           `json:"current_resin"`                 // 当前树脂
	MaxResin                  int           `json:"max_resin"`                     // 最大树脂
	ResinRecoveryTime         string        `json:"resin_recovery_time"`           // 树脂剩余恢复时间(s)
	FinishedTaskNum           int           `json:"finished_task_num"`             // 已完成委托
	TotalTaskNum              int           `json:"total_task_num"`                // 委托上限
	IsExtraTaskRewardReceived bool          `json:"is_extra_task_reward_received"` // 获得委托奖励
	RemainResinDiscountNum    int           `json:"remain_resin_discount_num"`     // 剩余周本减免
	ResinDiscountNumLimit     int           `json:"resin_discount_num_limit"`      // 总周本减免
	CurrentExpeditionNum      int           `json:"current_expedition_num"`        // 已完成派遣
	MaxExpeditionNum          int           `json:"max_expedition_num"`            // 派遣上限
	Expeditions               []Expeditions `json:"expeditions"`                   // 派遣数据
	CurrentHomeCoin           int           `json:"current_home_coin"`             // 已积累洞天宝钱
	MaxHomeCoin               int           `json:"max_home_coin"`                 // 洞天宝钱上限
	HomeCoinRecoveryTime      string        `json:"home_coin_recovery_time"`       // 洞天宝钱剩余恢复时间(s)
	CalendarUrl               string        `json:"calendar_url"`                  // 日历Url
	Transformer               Transformer   `json:"transformer"`                   // 质量参变仪数据
}

type Expeditions struct {
	AvatarSideIcon string `json:"avatar_side_icon"` // 委托角色头像
	Status         string `json:"status"`           // 状态
	RemainedTime   string `json:"remained_time"`    // 委托剩余时间(s)
}

type Transformer struct {
	Obtained     bool         `json:"obtained"`      // 是否拥有
	RecoveryTime RecoveryTime `json:"recovery_time"` // 恢复时间
	Wiki         string       `json:"wiki"`          // Wiki Url
	Noticed      bool         `json:"noticed"`       // 不知道是啥
	LatestJobId  string       `json:"latest_job_id"` // 不知道是啥
}

type RecoveryTime struct {
	Day     int  `json:"Day"`
	Hour    int  `json:"Hour"`
	Minute  int  `json:"Minute"`
	Second  int  `json:"Second"`
	Reached bool `json:"reached"` // 是否可用
}

type Sign struct {
	Status
	Data SignData `json:"data"`
}

type SignData struct {
	Code string `json:"code"`
}

type SignInfo struct {
	Status
	Data SignInfoData `json:"data"`
}

type SignInfoData struct {
	TotalSignDay  int    `json:"total_sign_day"`  // 总签到天数
	Today         string `json:"today"`           // 当天日期
	IsSign        bool   `json:"is_sign"`         // 已签到
	FirstBind     bool   `json:"first_bind"`      // 第一次绑定
	IsSub         bool   `json:"is_sub"`          // 不知道是啥
	MonthFirst    bool   `json:"month_first"`     // 当月首次签到
	SignCntMissed int    `json:"sign_cnt_missed"` // 本月漏签
}

type SignList struct {
	Status
	Data SignListData `json:"data"`
}

type SignListData struct {
	Month  int                `json:"month"`
	Awards []SignListDataItem `json:"awards"`
	Resign bool               `json:"resign"`
}

type SignListDataItem struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	Cnt  int    `json:"cnt"`
}
