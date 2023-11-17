package bot

// Self 机器人自身标识
type Self struct {
	Platform string `json:"platform"`
	UserId   string `json:"user_id"`
}

// Event 事件
type Event struct {
	Id         string  `json:"id"`
	Time       float64 `json:"time"`
	Type       string  `json:"type"`
	DetailType string  `json:"detail_type"`
	SubType    string  `json:"sub_type"`
}

// ActionReq 动作请求
type ActionReq struct {
	Action string         `json:"action"`
	Params map[string]any `json:"params"`
	Echo   string         `json:"echo,omitempty"`
	Self   Self           `json:"self,omitempty"`
}

// ActionResp 动作响应
type ActionResp struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Echo    string `json:"echo,omitempty"`
}

// MetaConnect 连接
type MetaConnect struct {
	Event
	Version struct {
		Impl          string `json:"impl"`
		Version       string `json:"version"`
		OnebotVersion string `json:"onebot_version"`
	} `json:"version"`
}

// MetaHeartbeat 心跳
type MetaHeartbeat struct {
	Event
	Interval int `json:"interval"`
}

// MetaStatusUpdate 状态更新
type MetaStatusUpdate struct {
	Event
	Status struct {
		Good bool `json:"good"`
		Bots []struct {
			Self struct {
				Platform string `json:"platform"`
				UserId   string `json:"user_id"`
			} `json:"self"`
			Online   bool   `json:"online"`
			QQStatus string `json:"qq.status,omitempty"`
		} `json:"bots"`
	} `json:"status"`
}

type Message struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

// MessagePrivate 私聊消息
type MessagePrivate struct {
	Event
	MessageId  string    `json:"message_id"`
	Message    []Message `json:"message"`
	AltMessage string    `json:"alt_message"`
	UserId     string    `json:"user_id"`
}

// MessageGroup 群消息
type MessageGroup struct {
	Event
	MessageId  string    `json:"message_id"`
	Message    []Message `json:"message"`
	AltMessage string    `json:"alt_message"`
	GroupId    string    `json:"group_id"`
	UserId     string    `json:"user_id"`
}

// MessageChannel 频道消息
type MessageChannel struct {
	Event
	MessageId  string    `json:"message_id"`
	Message    []Message `json:"message"`
	AltMessage string    `json:"alt_message"`
	GuildId    string    `json:"guild_id"`
	ChannelId  string    `json:"channel_id"`
	UserId     string    `json:"user_id"`
}

// NoticeFriend 好友增加/减少
type NoticeFriend struct {
	Event
	UserId string `json:"user_id"`
}

// NoticePrivateMessageDelete 私聊消息删除
type NoticePrivateMessageDelete struct {
	Event
	MessageId string `json:"message_id"`
	UserId    string `json:"user_id"`
}

// GetSelfInfo 获取机器人自身信息
type GetSelfInfo struct {
	ActionResp
	Data struct {
		UserId          string `json:"user_id"`
		UserName        string `json:"user_name"`
		UserDisplayname string `json:"user_displayname"`
	} `json:"data"`
}

// GetUserInfo 获取用户信息
type GetUserInfo struct {
	ActionResp
	Data struct {
		UserId          string `json:"user_id"`
		UserName        string `json:"user_name"`
		UserDisplayname string `json:"user_displayname"`
		UserRemark      string `json:"user_remark"`
	} `json:"data"`
}

// GetFriendList 获取好友列表
type GetFriendList struct {
	ActionResp
	Data []struct {
		UserId          string `json:"user_id"`
		UserName        string `json:"user_name"`
		UserDisplayname string `json:"user_displayname"`
		UserRemark      string `json:"user_remark"`
	} `json:"data"`
}

// NoticeGroupMember 群成员增加/减少
type NoticeGroupMember struct {
	Event
	UserId     string `json:"user_id"`
	GroupId    string `json:"group_id"`
	OperatorId string `json:"operator_id"`
}

// NoticeGroupMemberDelete 群消息删除
type NoticeGroupMemberDelete struct {
	Event
	GroupId    string `json:"group_id"`
	MessageId  string `json:"message_id"`
	UserId     string `json:"user_id"`
	OperatorId string `json:"operator_id"`
}
