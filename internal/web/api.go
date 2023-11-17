package web

import (
	"github.com/go-resty/resty/v2"
	"jjbot/core/logger"
	"net/http"
	"time"
)

var (
	HttpProxy string
)

func Get(uri string, headers map[string]string, proxy ...bool) (*resty.Response, error) {
	client := resty.New()
	client.SetRetryCount(3).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() == http.StatusTooManyRequests || err != nil
			},
		).
		SetTimeout(time.Second * 10)
	if len(proxy) > 0 && proxy[0] {
		client.SetProxy(HttpProxy)
	}
	resp, err := client.R().
		//EnableTrace().
		SetHeaders(headers).
		Get(uri)
	if err != nil {
		logger.SugarLogger.Errorf("重试次数耗尽: %v", err)
	}
	return resp, err
}

func Post(uri string, body any, isJson bool, headers map[string]string, proxy ...bool) (*resty.Response, error) {
	client := resty.New()
	client.SetRetryCount(3).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() == http.StatusTooManyRequests || err != nil
			},
		).
		SetTimeout(time.Second * 10)
	if len(proxy) > 0 && proxy[0] {
		client.SetProxy(HttpProxy)
	}
	if headers == nil {
		headers = map[string]string{}
	}
	if isJson {
		//headers["Content-Type"] = "application/json"
		client.SetHeader("Content-Type", "application/json")
	} else {
		client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.R().
		//EnableTrace().
		SetHeaders(headers).
		SetBody(body).
		Post(uri)
	if err != nil {
		logger.SugarLogger.Errorf("重试次数耗尽: %v", err)
	}
	return resp, err
}
