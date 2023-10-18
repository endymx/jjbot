package ys

import (
	"encoding/json"
	"github.com/pkg/errors"
	"jjbot/common/web"
	"jjbot/service/logger"
)

type Status struct {
	RetCode int    `json:"retcode"`
	Message string `json:"message"`
}

func getData(url string, post bool, data string, header map[string]string) (string, error) {
	headers := map[string]string{
		"x-rpc-app_version": version,
		"User_Agent":        "Mozilla/5.0 (Linux; Android 10; MIX 2 Build/QKQ1.190825.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.101 Mobile Safari/537.36 miHoYoBBS/2.35.2",
		"x-rpc-client_type": "5",
		"Referer":           "https://webstatic.mihoyo.com/",
	}
	if header != nil {
		for k, v := range header {
			headers[k] = v
		}
	}

	var status Status
	var j string
	if post {
		j = web.Post(url, data, false, nil, headers)
	} else {
		j = web.Get(url, nil, headers)
	}
	_ = json.Unmarshal([]byte(j), &status)
	if status.RetCode != 0 {
		logger.SugarLogger.Debug(j)
	}

	if status.Message == "" {
		return "", errors.Errorf("返回数据错误")
	}
	switch status.RetCode {
	case 0:
		return j, nil
	case -100:
		return "", errors.Errorf("cookie错误或失效，请重新获取Cookie并绑定")
	case 1008:
		return "", errors.Errorf("用户UID与Cookie不匹配")
	case 10001:
		return "", errors.Errorf("Cookie已达到30人上限，请退出并重登后再次获取Cookie")
	case 10102:
		return "", errors.Errorf("用户已经设置了隐私，无法查询")
	case -5003:
		return "", errors.Errorf("当前账号已签到，请勿重复签到")
	case -10001:
		return "", errors.Errorf("系统错误，请联系开发者")
	default:
		return "", errors.Errorf("未知错误，返回数据代码无对应\n返回错误信息：%s", status.Message)
	}
}
