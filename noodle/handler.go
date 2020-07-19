package noodle

type IMsgHandler interface {
	OnNewConnect(connect *WsConnect) bool         //新的消息队列
	OnDelConnect(connect *WsConnect)              //消息队列关闭
	ParserMsg(data []byte) (*Message, error)      //解析Message
	HandlerFunc(connect *WsConnect, msg *Message) //消息获得处理函数
}
