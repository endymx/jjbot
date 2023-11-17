package bot

import (
	"github.com/lxzan/gws"
	"github.com/tidwall/sjson"
	"jjbot/core/logger"
	"time"
)

var api *gws.Conn

func send(action string, params any, echo string) {
	if api != nil {
		logger.SugarLogger.Errorf("QQ机器人未成功登录，无法发送")
		return
	}
	json, _ := sjson.Set("", "action", action)
	if params != nil {
		json, _ = sjson.Set(json, "params", params)
	}
	if echo != "" {
		json, _ = sjson.Set(json, "echo", echo)
	}
	err := api.WriteMessage(gws.OpcodeText, []byte(json))
	if err != nil {
		logger.SugarLogger.Errorf("发送给QQ机器人消息失败：%s", err)
	}
}

//##################################信息类###################################

// SendPrivateMsg
/*
 * @Description 发送私聊消息
 * @param int user_id QQ号
 * @param string msg 要发送的内容
 * @param boolean auto_escape 消息内容是否作为纯文本发送（即不解析 CQ 码），只在 message 字段是字符串时有效
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func SendPrivateMsg(userId int64, msg string, autoEscape bool) {
	m := map[string]any{}
	m["user_id"] = userId
	m["message"] = msg
	m["auto_escape"] = autoEscape
	send("send_private_msg", m, "MessageData")
}

// SendGroupMsg
/**
 * @Description 发送群消息
 * @auth endymx
 * @param int group_id 群号
 * @param string msg 要发送的内容
 * @param boolean auto_escape 消息内容是否作为纯文本发送（即不解析 CQ 码），只在 message 字段是字符串时有效
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func SendGroupMsg(groupId int64, msg string, autoEscape bool) {
	time.Sleep(time.Second)
	m := map[string]any{}
	m["group_id"] = groupId
	m["message"] = msg
	m["auto_escape"] = autoEscape
	send("send_group_msg", m, "MessageData")
}

// SendGroupForwardMsg
/**
 * sendGroupForwardMsg
 * @Description 获取合并转发消息
 * @auth endymx
 * @param int group_id 群号
 * @param string msg 要发送的内容
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func SendGroupForwardMsg(groupId int64, msg []map[string]any) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["messages"] = msg
	send("send_group_forward_msg", m, "MessageData")
}

//SendMsg
/**
 * @Description 发送消息
 * @auth endymx
 * @param message_type
 * @param int user_id QQ号
 * @param int group_id 群号
 * @param string msg 要发送的内容
 * @param boolean auto_escape 消息内容是否作为纯文本发送（即不解析 CQ 码），只在 message 字段是字符串时有效
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func SendMsg(messageType any, userId int64, groupId int64, msg string, autoEscape bool) {
	m := map[string]any{}
	m["message_type"] = messageType
	m["user_id"] = userId
	m["group_id"] = groupId
	m["message"] = msg
	m["auto_escape"] = autoEscape
	send("send_msg", m, "MessageData")
}

/**
 * deleteMsg
 * @Description 撤回消息
 * @auth endymx
 * @param int message_id 消息 ID
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func deleteMsg(messageId int32) {
	m := map[string]any{}
	m["message_id"] = messageId
	send("delete_msg", m, "")
}

/**
 * deleteMsg
 * @Description 获取消息
 * @auth endymx
 * @param int message_id 消息 ID
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getMsg(messageId int32) {
	m := map[string]any{}
	m["message_id"] = messageId
	send("get_msg", m, "GetMessageData")
}

/**
 * getForwardMsg
 * @Description 获取合并转发消息
 * @auth endymx
 * @param int id 合并转发 ID
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getForwardMsg(id int32) {
	m := map[string]any{}
	m["id"] = id
	send("get_forward_msg", m, "GetForwardMessageData")
}

//##################################群组管理类###################################
/**
 * 群组踢人
 *
 * @param int group_id 群号
 * @param int user_id 要踢的 QQ 号
 * @param boolean request 拒绝此人的加群请求
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupKick(groupId int64, userId int64, request bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["user_id"] = userId
	m["reject_add_request"] = request
	send("set_group_kick", m, "")
}

/**
 * 群组单人禁言
 *
 * @param int group_id 群号
 * @param int user_id 要踢的 QQ 号
 * @param boolean duration 禁言时长，单位秒，0 表示取消禁言
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func SetGroupBan(groupId int64, userId int64, duration int) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["user_id"] = userId
	m["duration"] = duration
	send("set_group_ban", m, "")
}

/**
 * 群组匿名用户禁言
 *
 * @param int group_id 群号
 * @param object anonymous 可选，要禁言的匿名用户对象（群消息上报的 anonymous 字段）
 * @param string flag 可选，要禁言的匿名用户的 flag（需从群消息上报的数据中获得）
 * @param boolean duration 禁言时长，单位秒，无法取消匿名用户禁言
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupABan(groupId int64, anonymous any, flag string, duration int) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["anonymous"] = anonymous
	m["flag"] = flag
	m["duration"] = duration
	send("set_group_anonymous_ban", m, "")
}

/**
 * 群组全员禁言
 *
 * @param int group_id 群号
 * @param boolean enable 是否禁言
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupAllBan(groupId int64, enable bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["enable"] = enable
	send("set_group_whole_ban", m, "")
}

/**
 * 群组设置管理员
 *
 * @param int64 group_id 群号
 * @param int64 user_id 要踢的 QQ 号
 * @param bool enable true 为设置，false 为取消
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupAdmin(groupId int64, userId int64, enable bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["user_id"] = userId
	m["enable"] = enable
	send("set_group_admin", m, "")
}

/**
 * 群组匿名
 *
 * @param int group_id 群号
 * @param boolean enable true 为设置，false 为取消
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupAnonymous(groupId int64, enable bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["enable"] = enable
	send("set_group_anonymous", m, "")
}

/**
 * 设置群名片（群备注）
 *
 * @param int group_id 群号
 * @param int user_id 要设置的 QQ 号
 * @param string card 群名片内容，不填或空字符串表示删除群名片
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupCard(groupId int64, userId int64, card string) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["user_id"] = userId
	m["card"] = card
	send("set_group_card", m, "")
}

/**
 * 设置群名
 *
 * @param int group_id 群号
 * @param string group_name 新群名
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func SetGroupName(groupId int64, groupName string) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["group_name"] = groupName
	send("set_group_name", m, "")
}

/**
 * 退出群组
 *
 * @param int group_id 群号
 * @param boolean is_dismiss 是否解散，如果登录号是群主，则仅在此项为 true 时能够解散
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupLeave(groupId int64, isDismiss bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["is_dismiss"] = isDismiss
	send("set_group_leave", m, "")
}

/**
 * 设置群组专属头衔
 *
 * @param int group_id 群号
 * @param int user_id 要设置的 QQ 号
 * @param string special_title 专属头衔，不填或空字符串表示删除专属头衔
 * @param int duration 专属头衔有效期，单位秒，-1 表示永久，不过此项似乎没有效果，可能是只有某些特殊的时间长度有效，有待测试
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupTitle(groupId int64, user_id int32, special_title string, duration int) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["user_id"] = user_id
	m["special_title"] = special_title
	m["duration"] = duration
	send("set_group_special_title", m, "")
}

//##################################杂项类###################################
/**
 * 发送好友赞
 *
 * @param int user_id 要设置的 QQ 号
 * @param string times 赞的次数，每个好友每天最多 10 次
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func sendLike(userId int64, times int) {
	m := map[string]any{}
	m["times"] = times
	m["user_id"] = userId
	send("send_like", m, "")
}

/**
 * 处理加好友请求
 *
 * @param string flag 加好友请求的 flag（需从上报的数据中获得）
 * @param bool approve 是否同意请求
 * @param string remark 添加后的好友备注（仅在同意时有效）
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setFriendAdd(flag string, approve bool, remark string) {
	m := map[string]any{}
	m["flag"] = flag
	m["approve"] = approve
	m["remark"] = remark
	send("set_friend_add_request", m, "")
}

/**
 * 处理加群请求／邀请
 *
 * @param string flag 加好友请求的 flag（需从上报的数据中获得）
 * @param string type add 或 invite，请求类型（需要和上报消息中的 sub_type 字段相符）
 * @param bool approve 是否同意请求
 * @param string reason 拒绝理由（仅在拒绝时有效）
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setGroupAdd(flag string, typea string, approve bool, reason string) {
	m := map[string]any{}
	m["flag"] = flag
	m["type"] = typea
	m["approve"] = approve
	m["reason"] = reason
	send("set_group_add_request", m, "")
}

/**
 * 获取登录号信息
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func GetLoginInfo() {
	send("get_login_info", nil, "InfoData")
}

/**
 * 获取陌生人信息
 *
 * @param int user_id QQ 号
 * @param boolean no_cache 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getStrangerInfo(userId int64, noCache bool) {
	m := map[string]any{}
	m["user_id"] = userId
	m["no_cache"] = noCache
	send("get_stranger_info", m, "")
}

/**
 * 获取好友列表
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getFriendList() {
	send("get_friend_list", nil, "StrangerInfoData")
}

/**
 * 获取群信息
 *
 * @param int group_id 群号
 * @param boolean no_cache 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getFriendInfo(groupId int64, noCache bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["no_cache"] = noCache
	send("get_group_info", m, "GroupData")
}

/**
 * 获取群列表
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getGroupList() {
	send("get_group_list", nil, "GroupList")
}

/**
 * 获取群成员信息
 *
 * @param int group_id 群号
 * @param int user_id QQ 号
 * @param boolean no_cache 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getGroupMemberInfo(groupId int64, userId int64, noCache bool) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["user_id"] = userId
	m["no_cache"] = noCache
	send("get_group_member_info", m, "GroupInfoData")
}

/**
 * 获取群成员列表
 *
 * @param int group_id 群号
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getGroupMemberList(groupId int64) {
	m := map[string]any{}
	m["group_id"] = groupId
	send("get_group_member_list", m, "MemberInfoData")
}

/**
 * 获取群荣誉信息
 *
 * @param int group_id 群号
 * @param string typea 要获取的群荣誉类型，可传入 talkative performer legend strong_newbie emotion 以分别获取单个类型的群荣誉数据，或传入 all 获取所有数据
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getGroupHonorInfo(groupId int64, typea string) {
	m := map[string]any{}
	m["group_id"] = groupId
	m["type"] = typea
	send("get_group_honor_info", m, "HonorInfoData")
}

/**
 * 获取 Cookies
 *
 * @param string domain 需要获取 cookies 的域名
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getCookies(domain string) {
	m := map[string]any{}
	m["domain"] = domain
	send("get_cookies", m, "")
}

/**
 * 获取 CSRF Token
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getCsrfToken(domain string) {
	m := map[string]any{}
	m["domain"] = domain
	send("get_csrf_token", m, "")
}

/**
 * 获取 QQ 相关接口凭证
 *
 * @param string domain 需要获取 cookies 的域名
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getCredentials(domain string) {
	m := map[string]any{}
	m["domain"] = domain
	send("get_credentials", m, "")
}

/**
 * 获取语音
 *
 * @param string file 收到的语音文件名（消息段的 file 参数），如 0B38145AA44505000B38145AA4450500.silk
 * @param string out_format 要转换到的格式，目前支持 mp3、amr、wma、m4a、spx、ogg、wav、flac
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getRecord(file string, outFormat string) {
	m := map[string]any{}
	m["file"] = file
	m["out_format"] = outFormat
	send("get_record", m, "")
}

/**
 * 获取图片
 *
 * @param string file 收到的图片文件名（消息段的 file 参数），如 6B4DE3DFD1BD271E3297859D41C530F5.jpg
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getImage(file string) {
	m := map[string]any{}
	m["file"] = file
	send("get_image", m, "")
}

/**
 * 检查是否可以发送图片
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func canSendImage() {
	send("can_send_image", nil, "")
}

/**
 * 检查是否可以发送语音
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func canSendRecord() {
	send("can_send_record", nil, "")
}

/**
 * 获取运行状态
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getStatus() {
	send("get_status", nil, "")
}

/**
 * 获取版本信息
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func getVersionInfo() {
	send("get_version_info", nil, "")
}

/**
 * 重启 OneBot 实现
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func setRestart() {
	send("set_restart", nil, "")
}

/**
 * 清理缓存
 *
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func cleanCache() {
	send("clean_cache", nil, "")
}

/**
 * 快速操作
 *
 * @param array context 事件数据对象，可做精简，如去掉 message 等无用字段
 * @param array operation 快速操作对象，例如 ["ban" => true, "reply" => "请不要说脏话"]
 * @return mixed (true|null|false)只要不返回false并且网络没有断开，而且服务端接收正常，数据基本上可以看做100%能发过去
 */
func quickOperation(context any, operation any) {
	m := map[string]any{}
	m["context"] = context
	m["operation"] = operation
	send(".handle_quick_operation", m, "")
}
