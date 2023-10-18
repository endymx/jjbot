package qq

type GroupMessage struct {
	Time        int64  `json:"time"`           // 事件发生的时间戳
	Date        string `json:"date,omitempty"` // 日期
	SelfId      int64  `json:"self_id"`        // 收到事件的机器人QQ号
	PostType    string `json:"post_type"`      // 上报类型 (可能是message)
	MessageType string `json:"message_type"`   // 消息类型 (可能是group)
	SubType     string `json:"sub_type"`       // 消息子类型，正常消息是normal，匿名消息是anonymous，系统提示是notice
	MessageId   int64  `json:"message_id"`     // 消息ID
	GroupId     int64  `json:"group_id"`       // 群号
	UserId      int64  `json:"user_id"`        // 发送者QQ号
	Message     string `json:"message"`        // 消息内容
	RawMessage  string `json:"raw_message"`    // 原始消息内容
	Sender      Sender // 发送人信息
}

type Sender struct {
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserId   int    `json:"user_id"`
}
