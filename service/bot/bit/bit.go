package bit

import (
	"context"
	Rand "crypto/rand"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"jjbot/common/db/mongodb"
	"jjbot/common/db/redis"
	"jjbot/service/bot/botapi"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var Max = 0.5
var Min = 0.0002
var tj = false
var Command = []bool{false, false, false, false}

func float8b(f float64) float64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", f+0.000000005), 64)
	return num
}

func Query(uid int64) Bit {
	var newBit Bit
	mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": uid}, nil, &newBit)
	return newBit
}

func AddHash(uid int64, hash float64) {
	var newBit Bit
	mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": uid}, nil, &newBit)
	newBit.Hash += hash
	mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": uid}, newBit)
	newBit = Bit{}
	mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": -1}, nil, &newBit)
	newBit.Hash += hash
	if newBit.Hash >= newBit.Bit {
		var max Bit
		mongodb.Client["endymx"].FindUnmarshal(
			"bot",
			"bit",
			bson.D{{"uid", bson.M{"$exists": true}}},
			nil,
			&Bit{},
			func(i int, a *any) {
				p := (*a).(*Bit)
				if p.Uid != -1 {
					f := float8b(p.Hash / newBit.Bit)
					p.Bit = float8b(p.Bit + f)
					p.Hash = 0
					if f > max.Bit {
						max.Uid = p.Uid
						max.Bit = f
					}
					mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": p.Uid}, p)
				}
			})
		rand.New(rand.NewSource(time.Now().UnixNano()))
		newBit.Bit = newBit.Bit * (float64(9500+rand.Intn(3500)) / 10000)
		newBit.Hash = 0
		botapi.SendGroupMsg(811635507,
			fmt.Sprintf("成功鸽了一次，1鸽子已被分发\n本次最佳鸽手为[CQ:at,qq=%d]，共获得%.8f鸽子", max.Uid, max.Bit), false)
	}
	mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": -1}, newBit)
}

func MoveBit(uid int64, uid2 int64, bit float64, messageId int64, gid int64) {
	var newBit, newBit2 Bit
	mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": uid}, nil, &newBit)
	if newBit.Uid == 0 || newBit.Bit <= bit {
		botapi.SendGroupMsg(gid,
			fmt.Sprintf("[CQ:reply,id=%d]你的鸽子余额不足，无法鸽别人", messageId), false)
		return
	}
	mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": uid2}, nil, &newBit2)
	if newBit2.Uid == 0 {
		botapi.SendGroupMsg(gid,
			fmt.Sprintf("[CQ:reply,id=%d]对方尚未搭建鸽巢", messageId), false)
		return
	}
	newBit.Bit = float8b(newBit.Bit - bit)
	newBit2.Bit = float8b(newBit.Bit + bit)
	mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": uid}, newBit)
	mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": uid2}, newBit2)
	botapi.SendGroupMsg(gid,
		fmt.Sprintf(
			"%.8f鸽子离开了他的主人\n[CQ:at,qq=%d]：%.8f -> %.8f\n[CQ:at,qq=%d]：%.8f -> %.8f",
			bit,
			uid,
			newBit.Bit+bit,
			newBit.Bit,
			uid2,
			newBit2.Bit-bit,
			newBit2.Bit,
		), false)
}

func UpdateFood() {
	t, _ := strconv.ParseInt(time.Now().Format("15"), 10, 64)
	ts, _ := strconv.ParseInt(time.Now().Format("04"), 10, 64)
	if t >= 1 && t < 8 { //凌晨2~6点不播报
		return
	}
	var newBit Bit
	mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": -1}, nil, &newBit)
	if newBit.Food == Max || newBit.Food == Min {
		if t == 0 || t == 12 {
			mongodb.Client["endymx"].FindUnmarshal(
				"bot",
				"bit",
				bson.D{{"uid", bson.M{"$exists": true}}},
				nil,
				&Bit{},
				func(i int, a *any) {
					p := (*a).(*Bit)
					if p.Uid != -1 {
						p.Food = 0
						mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": p.Uid}, p)
					}
				})
			newBit.Food = 0.01
			tj = false
		} else {
			return
		}
	}
	n, _ := Rand.Int(Rand.Reader, big.NewInt(800000))
	rand.ExpFloat64()
	r := float64(700000+n.Int64()) / 1000000
	n, _ = Rand.Int(Rand.Reader, big.NewInt(9))
	if n.Int64() <= 6 {
		n, _ = Rand.Int(Rand.Reader, big.NewInt(400000))
		r = float64(800000+n.Int64()) / 1000000
	}
	newBit.Food = float8b(newBit.Food * r)
	if newBit.Food > Max {
		newBit.Food = Max
	} else if newBit.Food < Min {
		newBit.Food = Min
	}
	mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": -1}, newBit)
	if newBit.Food == Max || newBit.Food == Min {
		if tj {
			return
		}
		if ts != 30 && ts != 0 {
			return
		}
		botapi.SendGroupMsg(811635507,
			"鸽粮快报！当前鸽粮被价格管理局处罚，请购买者尽快联系本店退回！\n（仅可卖出不可买入，将在凌晨0点清空所有鸽粮并重置价格）", false)
		botapi.SendGroupMsg(198848645,
			"鸽粮快报！当前鸽粮被价格管理局处罚，请购买者尽快联系本店退回！\n（仅可卖出不可买入，将在凌晨或中午12点清空所有鸽粮并重置价格）", false)
		tj = true
	} else if newBit.Food < Max || newBit.Food > Min {
		if tj {
			return
		}
		if ts != 30 && ts != 0 {
			return
		}
		botapi.SendGroupMsg(811635507,
			fmt.Sprintf("鸽粮快报！当前鸽粮价格为%.8f鸽子！", newBit.Food), false)
		botapi.SendGroupMsg(198848645,
			fmt.Sprintf("鸽粮快报！当前鸽粮价格为%.8f鸽子！", newBit.Food), false)
	}
}

func RP(rawMessage string, messageId int64, groupId int64, userId int64) {
	rp := RedPacket{}
	ctx := context.Background()
	_ = redis.Client.HGetAll(ctx, "redpacket").Scan(&rp)
	if rp.Status {
		if strings.TrimSpace(rawMessage) != "鸽鸽" {
			botapi.SendGroupMsg(groupId,
				fmt.Sprintf("[CQ:reply,id=%d]当前正在派发，请等待领取完成后重新派发", messageId), false)
			return
		}
		rpn, _ := redis.Client.HGetAll(ctx, "redpacket_number").Result()
		for n := range rpn {
			if n == strconv.FormatInt(userId, 10) {
				return
			}
		}
		if rp.Share > 1 {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			rpn[strconv.FormatInt(userId, 10)] = fmt.Sprintf("%.8f", rp.Total/2*rand.Float64())
		} else {
			rpn[strconv.FormatInt(userId, 10)] = fmt.Sprintf("%.8f", rp.Total)
		}
		n, _ := strconv.ParseFloat(rpn[strconv.FormatInt(userId, 10)], 64)
		rp.Total -= n
		rp.Share--
		if rp.Share == 0 {
			s := ""
			max := []float64{0, 0}
			for k, v := range rpn {
				uid, _ := strconv.ParseInt(k, 10, 64)
				git, _ := strconv.ParseFloat(v, 64)
				if git > max[1] {
					max = []float64{float64(uid), git}
				}
				s = fmt.Sprintf("%s\n%d 获得 %.8f 鸽子", s, uid, git)
				var newBit Bit
				mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": uid}, nil, &newBit)
				if newBit.Uid == 0 {
					s = fmt.Sprintf("%s，但是没有鸽窝无法领取", s)
					var newBit2 Bit
					mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": rp.Uid}, nil, &newBit2)
					newBit2.Bit += git
					mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": rp.Uid}, newBit2)
					continue
				}
				newBit.Bit += git
				mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": userId}, newBit)
			}
			f := []map[string]any{{
				"type": "node",
				"data": map[string]any{"name": "鸽鸽", "uin": "1877035391", "content": fmt.Sprintf("鸽子被领取完毕，最佳鸽手为 %0.f", max[0])},
			},
				{
					"type": "node",
					"data": map[string]any{"name": "鸽鸽", "uin": "1877035391", "content": fmt.Sprintf("鸽子分发如下：%s", s)},
				}}
			botapi.SendGroupForwardMsg(groupId, f)
			rp = RedPacket{}
			redis.Client.HSet(ctx, "redpacket", rp)
			redis.Client.Del(context.Background(), "redpacket_number")
			return
		}
		redis.Client.HSet(ctx, "redpacket_number", rpn)
	} else {
		var newBit Bit
		mongodb.Client["endymx"].FindOneUnmarshal("bot", "bit", bson.M{"uid": userId}, nil, &newBit)
		m := strings.Split(rawMessage, " ")
		if len(m) < 3 {
			botapi.SendGroupMsg(groupId,
				fmt.Sprintf("[CQ:reply,id=%d]当前没有正在派发的鸽子，请使用\"鸽鸽 [总额] [份数]\"来新建派送点", messageId), false)
			return
		}
		git, _ := strconv.ParseFloat(m[1], 64)
		if git < 0.00001 {
			botapi.SendGroupMsg(groupId,
				fmt.Sprintf("[CQ:reply,id=%d]鸽子数值错误，请填入大于0.00001的数值", messageId), false)
			return
		}
		if newBit.Bit < git {
			botapi.SendGroupMsg(groupId,
				fmt.Sprintf("[CQ:reply,id=%d]你的鸽子余额不足，无法创建", messageId), false)
			return
		}
		share, _ := strconv.ParseInt(m[2], 10, 64)
		if share <= 0 {
			botapi.SendGroupMsg(groupId,
				fmt.Sprintf("[CQ:reply,id=%d]数量错误，请填入大于0的数值", messageId), false)
			return
		}
		newBit.Bit -= git
		mongodb.Client["endymx"].UpdateOne("bot", "bit", bson.M{"uid": userId}, newBit)
		rp = RedPacket{
			Status: true,
			Uid:    userId,
			Total:  git,
			Share:  share,
		}
		botapi.SendGroupMsg(groupId,
			fmt.Sprintf("[CQ:at,qq=%d]创建了新的鸽子派送点（含%.8f鸽子，共%d份），快输入\"鸽鸽\"来抢鸽子吧", userId, git, share), false)
		redis.Client.HSet(ctx, "redpacket", rp)
	}
}
