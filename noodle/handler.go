package noodle

type HandlerFunc func(connect *WsConnect, msg *Message) bool
type IMsgHandler interface {
	OnNewConnect(connect *WsConnect) bool                        //新的消息队列
	OnDelConnect(connect *WsConnect)                             //消息队列关闭
	OnProcessMsg(connect *WsConnect, msg *Message) bool          //默认的消息处理函数
	GetHandlerFunc(connect *WsConnect, msg *Message) HandlerFunc //根据消息获得处理函数
}
