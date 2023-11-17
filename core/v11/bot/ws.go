package bot

import (
	"github.com/tidwall/gjson"
	"jjbot/core/logger"
)

func ws(addr string, path string) {
	json := ""
	if path == "/api" {
		if gjson.Get(json, "status").String() != "ok" {
			logger.SugarLogger.Infof("API报错: %s", json)
			return
		}
		switch gjson.Get(json, "echo").String() {
		case "MessageData": //send_private_msg()回调 \ send_group_msg()回调 \ send_msg()回调
			//do something
			break
		case "GetMessageData": //get_msg()回调
			//do something
			break
		case "GetForwardMessageData": //get_forward_msg()回调
			//do something
			break
		case "InfoData": //getLoginInfo()回调
			logger.SugarLogger.Infof("登录QQ账号: %s (%d)",
				gjson.Get(json, "data.nickname").String(), gjson.Get(json, "data.user_id").Int())
			//do something
			break
		case "StrangerInfoData": //get_stranger_info()回调
			//do something
			break
		case "FriendData": //get_friend_list()回调
			//do something
			break
		case "GroupData": //get_group_info()回调
			//do something
			break
		case "GroupList": //get_group_list()回调
			//do something
			break
		case "GroupInfoData": //get_group_info()回调
			//do something
			break
		case "MemberInfoData": //get_group_member_info()回调 \ get_group_member_list() array <>回调
			//do something
			break
		case "HonorInfoData": //get_group_honor_info()回调
			//do something
			break
		case "ImageInfoData": //get_image()回调
			//do something
			break
		case "RecordInfoData": //get_record()回调
			//do something
			break
		case "CanSendImageData": //can_send_image()回调
			//do something
			break
		case "CanSendRecordData": //can_send_record()回调
			//do something
			break
		case "PluginStatusData": //get_status() <PluginsGoodData>回调
			//do something
			break
		/*
			case "PluginsGoodData"://get_status()回调
				//do something
				break;
		*/
		case "VersionInfoData": //get_version_info()回调
			//do something
			break
		default:
			logger.SugarLogger.Infof("无法识别的API返回: %s", json)
			break
		}
		onMessageApi(json)
	} else {
		switch gjson.Get(json, "post_type").String() {
		case "message": //消息事件
			if gjson.Get(json, "message_type").String() == "private" {
				go onPrivateMessageEvent(json, gjson.Get(json, "time").Int(),
					gjson.Get(json, "self_id").Int(), gjson.Get(json, "post_type").String(),
					gjson.Get(json, "message_type").String(), gjson.Get(json, "sub_type").String(),
					gjson.Get(json, "message_id").Int(), gjson.Get(json, "user_id").Int(),
					gjson.Get(json, "message").String(), gjson.Get(json, "raw_message").String(),
					gjson.Get(json, "font").Int(), gjson.Get(json, "sender").Map())
			} else if gjson.Get(json, "message_type").String() == "group" {
				go onGroupMessageEvent(json, gjson.Get(json, "time").Int(),
					gjson.Get(json, "self_id").Int(), gjson.Get(json, "post_type").String(),
					gjson.Get(json, "message_type").String(), gjson.Get(json, "sub_type").String(),
					gjson.Get(json, "message_id").Int(), gjson.Get(json, "group_id").Int(), gjson.Get(json, "user_id").Int(),
					gjson.Get(json, "message").String(), gjson.Get(json, "raw_message").String(),
					gjson.Get(json, "font").Int(), gjson.Get(json, "sender").Map())
			}
			break
		case "notice": //通知事件
			switch gjson.Get(json, "notice_type").String() {
			case "group_upload": //群文件上传
				//GroupFileUploadEvent();
				break
			case "group_admin": //群管理员变动
				//GroupAdminEvent($data["time"], $data["self_id"], $data["post_type"], $data["notice_type"],
				//$data["sub_type"], $data["group_id"], $data["user_id"])
				break
			case "group_decrease": //群成员减少
				//do something
				break
			case "group_increase": //群成员增加
				//do something
				break
			case "group_ban": //群禁言
				//GroupBanEvent($data["time"], $data["self_id"], $data["post_type"], $data["notice_type"],
				//$data["sub_type"], $data["operator_id"], $data["group_id"], $data["user_id"], $data["duration"])
				break
			case "friend_add": //好友添加
				//FriendAddEvent($data["time"], $data["self_id"], $data["post_type"], $data["notice_type"],
				//$data["user_id"])
				break
			case "group_recall": //群消息撤回
				//do something
				break
			case "friend_recall": //好友消息撤回
				//do something
				break
			case "notify": //群通知
				switch gjson.Get(json, "sub_type").String() {
				case "poke": //戳一戳
					//do something
					break
				case "lucky_king": //红包运气王
					//do something
					break
				case "honor": //群荣誉
					//do something
					break
				}
				break
			}
			break
		case "request": //请求事件
			if gjson.Get(json, "request_type").String() == "friend" { //加好友请求
				//FriendRequestEvent($data, $data["time"], $data["self_id"], $data["post_type"], $data["request_type"], $data["user_id"], $data["comment"], $data["flag"])
			} else if gjson.Get(json, "request_type").String() == "group" { //加群请求or邀请
				onGroupMemberAddRequestEvent(json, gjson.Get(json, "time").Int(),
					gjson.Get(json, "self_id").Int(), gjson.Get(json, "post_type").String(),
					gjson.Get(json, "request_type").String(), gjson.Get(json, "sub_type").String(),
					gjson.Get(json, "group_id").Int(), gjson.Get(json, "user_id").Int(),
					gjson.Get(json, "comment").String(), gjson.Get(json, "flag").String())
			}
			break
		case "meta_event": //元事件

			break
		case "CQLifecycleMetaEvent": //ws生命周期
			//do something
			break
		case "CQHeartbeatMetaEvent": //ws心跳
			//do something
			break
		default:
			logger.SugarLogger.Infof("无法识别的Event返回: %s", json)
			break
		}
	}
}
