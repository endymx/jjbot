package ys

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"jjbot/common/db/mongodb"
	"jjbot/service/logger"
	"strconv"
	"time"
)

var (
	version = "2.11.1"
	month   int
	item    []string
)

type User struct {
	QQ      int64
	Group   int64
	Discord string
	DGroup  string
	Uid     int64
	Cookie  string
	Date    string
}

func BindUid(user, data User) {
	if user.Uid == 0 {
		mongodb.Client["endymx"].InsertOne("bot", "ys", data)
	} else {
		data.Group = user.Group
		mongodb.Client["endymx"].UpdateOne("bot", "ys", user, data)
	}
}

func BindGroup(user, data User) {
	mongodb.Client["endymx"].UpdateOne("bot", "ys", user, data)
}

func GetDaily(user User) (*Daily, error) {
	headers := map[string]string{
		"Cookie": user.Cookie,
		"DS":     getDS(fmt.Sprintf("role_id=%d&server=%s", user.Uid, getServer(user.Uid)), nil),
	}

	s := Daily{}
	d, err := getData(fmt.Sprintf("https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/dailyNote?server=%s&role_id=%d", getServer(user.Uid), user.Uid), false, "", headers)
	_ = json.Unmarshal([]byte(d), &s)
	s.Data.ResinRecoveryTime = resolveTime(s.Data.ResinRecoveryTime)
	return &s, err
}

func GetSign(user User) (*Sign, error) {
	data, _ := sjson.Set("", "act_id", "e202009291139501")
	data, _ = sjson.Set(data, "uid", user.Uid)
	data, _ = sjson.Set(data, "region", getServer(user.Uid))

	headers := map[string]string{
		"User_Agent":        "Mozilla/5.0 (Linux; Android 10; MIX 2 Build/QKQ1.190825.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.101 Mobile Safari/537.36 miHoYoBBS/2.3.0",
		"Cookie":            user.Cookie,
		"x-rpc-device_id":   getRandomHex(32),
		"Origin":            "https://webstatic.mihoyo.com",
		"X_Requested_With":  "com.mihoyo.hyperion",
		"DS":                getOldDS(true),
		"x-rpc-client_type": "2",
		"Referer":           "https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=e202009291139501&utm_source=bbs&utm_medium=mys&utm_campaign=icon",
		"x-rpc-app_version": "2.28.1",
	}

	s := Sign{}
	d, err := getData("https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign", true, data, headers)
	_ = json.Unmarshal([]byte(d), &s)
	return &s, err
}

func GetSignInfo(user User) (*SignInfo, error) {
	headers := map[string]string{"Cookie": user.Cookie}

	s := SignInfo{}
	d, err := getData(fmt.Sprintf("https://api-takumi.mihoyo.com/event/bbs_sign_reward/info?act_id=e202009291139501&region=%s&uid=%d", getServer(user.Uid), user.Uid), false, "", headers)
	_ = json.Unmarshal([]byte(d), &s)
	return &s, err
}

func GetSignList() ([]string, error) {
	m, _ := strconv.Atoi(time.Now().Format("01"))
	if month != m {
		data, err := getData("https://api-takumi.mihoyo.com/event/bbs_sign_reward/home?act_id=e202009291139501", false, "", nil)
		if err == nil {
			s := SignList{}
			_ = json.Unmarshal([]byte(data), &s)

			item = []string{}
			for _, i := range s.Data.Awards {
				item = append(item, fmt.Sprintf("%s*%d", i.Name, i.Cnt))
			}
			month = int(gjson.Get(data, "data.month").Int())
		} else {
			logger.SugarLogger.Error(err)
			return nil, err
		}
	}
	return item, nil
}

func DailyMsg(s *Daily) string {
	msg := fmt.Sprintf("当前树脂：%d/%d", s.Data.CurrentResin, s.Data.MaxResin)
	if s.Data.CurrentResin < s.Data.MaxResin {
		msg = fmt.Sprintf("%s\n树脂溢满：%s", msg, s.Data.ResinRecoveryTime)
	}
	msg = fmt.Sprintf("%s\n每日任务：%d/%d", msg, s.Data.FinishedTaskNum, s.Data.TotalTaskNum)
	if s.Data.IsExtraTaskRewardReceived {
		msg = fmt.Sprintf("%s(已领取奖励)", msg)
	} else {
		msg = fmt.Sprintf("%s(未领取奖励)", msg)
	}
	msg = fmt.Sprintf("%s\n洞天宝钱积累：%d/%d", msg, s.Data.CurrentHomeCoin, s.Data.MaxHomeCoin)
	msg = fmt.Sprintf("%s\n周本半价剩余：%d", msg, s.Data.RemainResinDiscountNum)
	if s.Data.Transformer.Obtained && s.Data.Transformer.RecoveryTime.Reached {
		msg = fmt.Sprintf("%s\n质量参变仪已可用", msg)
	}
	return msg
}
