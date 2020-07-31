package noodle

import "noodlenet/noodle/msg"

type HandlerFunc func(connect *WsConnect, msg *msg.Pb) bool

type IMsgHandler interface {
	OnNewConnect(connect *WsConnect) bool           //新的消息队列
	OnDelConnect(connect *WsConnect)                //消息队列关闭
	//MessageFromBytes(data []byte) (*Message, error) //解析Message
	//MessageToBytes(msg *Message) ([]byte, error)    //Message转数据
	HandlerFunc(connect *WsConnect, msg *msg.Pb)   //消息获得处理函数
}
