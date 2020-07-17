package noodle

type HandlerFunc func(msgque *WsConnect, msg *Message) bool
type IMsgHandler interface {
	OnNewMsgQue(msgque *WsConnect) bool                //新的消息队列
	OnDelMsgQue(msgque *WsConnect)                     //消息队列关闭
	OnProcessMsg(msgque *WsConnect, msg *Message) bool //默认的消息处理函数
	//OnConnectComplete(msgque  *WsConnect, ok bool) bool                  //连接成功
	GetHandlerFunc(msgque *WsConnect, msg *Message) HandlerFunc //根据消息获得处理函数
}
