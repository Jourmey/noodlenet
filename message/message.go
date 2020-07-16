package message

type MessageHead struct {
	Len   uint32 //数据长度
	Error uint16 //错误码
	Cmd   uint8  //命令
	Act   uint8  //动作
	Index uint16 //序号
	Flags uint16 //标记

	forever bool
	data    []byte
}

type Message struct {
	Head *MessageHead //消息头，可能为nil
	Data []byte       //消息数据
	User interface{}  //用户自定义数据
}
