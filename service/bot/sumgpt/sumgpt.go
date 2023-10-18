package sumgpt

import (
	"encoding/json"
	"fmt"
	"jjbot/common/db/mongodb"
	"jjbot/common/web"
	"strings"
	"time"
)

const TokenThreshold = 1200

var Chat = ""
var Token = 0
var Click = false
var Sums []map[string]any
var LastSumTime = time.Now().Format("2006-01-02 15:04:05")
var ChatCount = 0

func Sum() {
	now := time.Now().Format("2006-01-02 15:04:05")
	// calculate time after last sum, if less than 10 minutes, return
	if LastSumTime != "" {
		lastSumTime, _ := time.Parse("2006-01-02 15:04:05", LastSumTime)
		if time.Now().Sub(lastSumTime).Minutes() < 10 {
			return
		}
	}

	h := map[string]string{"Authorization": "Bearer SUMGPT_TOKEN", "content-type": "text/plain"}
	msg := web.Post("https://sumgpt.jp.nico.wang:44443/sum", Chat, false, nil, h)
	Chat = ""
	Token = 0
	s := API{}
	_ = json.Unmarshal([]byte(msg), &s)
	var resultMessage string
	resultMessage = fmt.Sprintf("总结%s至%s时间段的聊天记录如下：\n%s", LastSumTime, now, strings.ReplaceAll(s.RawText, "\r\n", "\n"))
	if s.Code != 0 {
		resultMessage = "错误：\n" + s.Message
	}
	newSums := []map[string]any{{"type": "node", "data": map[string]any{"name": "SumGPT", "uin": "932141154", "content": resultMessage}}}
	mongodb.Client["endymx"].InsertOne("bot", "sumgpt-sums", newSums[0])
	Sums = append(newSums, Sums...)
	if len(Sums) >= 5 {
		Sums = Sums[:5]
	}
	LastSumTime = now
	return
}
