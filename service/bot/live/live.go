package live

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"go.mongodb.org/mongo-driver/bson"
	"jjbot/common/db/mongodb"
	"jjbot/common/db/redis"
	"jjbot/common/web"
	"jjbot/service/bot/botapi"
)

type Data struct { //所有首字母一定要大写以导出，不然读不到数据会导致只有_id
	User     int64
	Group    int64
	LiveChat bool
}

func GetLiveInfo() {
	db := mongodb.Client["endymx"].Find("bot", "live", bson.D{{"user", bson.M{"$exists": true}}}, nil)
	if db != nil {
		ctx := context.Background()
		json := ""
		for i, d := range db {
			json, _ = sjson.Set(json, fmt.Sprintf("uids.%d", i), d.Map()["user"])
		}
		data := web.Post("https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids", json, true, nil, nil)
		gjson.Get(data, "data").ForEach(func(uid, user gjson.Result) bool {
			if user.Get("live_status").Int() == 1 {
				b, _ := redis.Client.HGet(ctx, "live", uid.String()).Bool()
				if !b {
					redis.Client.HSet(ctx, "live", uid.String(), true)
					name := user.Get("uname").String()
					if name == "红米煮" {
						name = "红米猪"
					}
					MassLive(db, uid.Int(), fmt.Sprintf(
						"[CQ:image,file=%s,subType=1]\n"+
							"主播 %s 开播啦\n"+
							"https://live.bilibili.com/%d",
						user.Get("cover_from_user").String(),
						name,
						user.Get("room_id").Int(),
					), true)
				}
			} else {
				b, _ := redis.Client.HGet(ctx, "live", uid.String()).Bool()
				if b {
					redis.Client.HDel(ctx, "live", uid.String())
					name := user.Get("uname").String()
					if name == "红米煮" {
						name = "红米猪"
					}
					MassLive(db, uid.Int(), fmt.Sprintf(`主播 %s 跑路啦!`, name), true)
				}
				b, _ = redis.Client.HGet(ctx, "livechat", uid.String()).Bool()
				if b {
					redis.Client.HDel(ctx, "live", uid.String())
					//关闭livechat
				}
			}
			return true
		})
	}
}

func MassLive(db []bson.D, uid int64, message string, of bool) {
	for _, d := range db {
		user := d.Map()["user"].(int64)
		if of {
			if d.Map()["livechat"] == true { //TODO 需要添加 && LiveChat聊天室部分未启动的条件
				//go启动livechat
			}
		}
		if uid == user {
			botapi.SendGroupMsg(d.Map()["group"].(int64), message, false)
		}
	}
}

func AddLive(uid int, gid int64, chat bool) {
	json, _ := sjson.Set("", "uids.0", uid)
	data := web.Post("https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids", json, true, nil, nil)
	if !gjson.Get(data, fmt.Sprintf("data.%d", uid)).Exists() {
		botapi.SendGroupMsg(gid, "未找到用户，请检查用户id", false)
		return
	}
	if uid == 388832538 {
		data, _ = sjson.Set(data, fmt.Sprintf("data.%d.uname", uid), "红米猪")
	}
	liveData := Data{
		User:     int64(uid),
		Group:    gid,
		LiveChat: chat,
	}
	if mongodb.Client["endymx"].FindOne("bot", "live", bson.D{{"user", uid}, {"group", gid}}, nil) == nil {
		mongodb.Client["endymx"].InsertOne("bot", "live", liveData)
	} else {
		mongodb.Client["endymx"].UpdateOne("bot", "live",
			bson.D{{"user", uid}, {"group", gid}},
			liveData,
		)
	}
	botapi.SendGroupMsg(gid, "订阅主播 "+gjson.Get(data, fmt.Sprintf("data.%d.uname", uid)).String(), false)
}

func RemoveLive(uid int, gid int64) {
	if mongodb.Client["endymx"].FindOne("bot", "live", bson.D{{"user", uid}, {"group", gid}}, nil) != nil {
		json, _ := sjson.Set("", "uids.0", uid)
		data := web.Post("https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids", json, true, nil, nil)
		if uid == 388832538 {
			data, _ = sjson.Set(data, fmt.Sprintf("data.%d.uname", uid), "红米猪")
		}
		mongodb.Client["endymx"].DeleteOne("bot", "live", bson.D{{"user", uid}, {"group", gid}})
		botapi.SendGroupMsg(gid, "取消订阅主播 "+gjson.Get(data, fmt.Sprintf("data.%d.uname", uid)).String(), false)
	} else {
		botapi.SendGroupMsg(gid, "本群未订阅过该主播", false)
	}
}
