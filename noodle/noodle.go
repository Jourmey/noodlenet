package noodle

import (
	"sync"
)

var ConnectId uint32 //消息队列id
var ConnectMapSync sync.Mutex
var ConnectMap = map[uint32]*WsConnect{}

//var GMManager GMessageManager //管理广播信息

type gMessage struct {
	c   chan struct{} // 触发广播使用
	msg *Message
	fun func(connect *WsConnect) bool
}

type GMessageManager struct {
	gMessageArray [65535]*gMessage
	gId           uint16
}

func (g *GMessageManager) GetGMessage() *gMessage {
	return g.gMessageArray[g.gId]
}

type ConnType uint8

const (
	ConnTypeListen ConnType = iota //监听
	ConnTypeConn                   //连接产生的
	ConnTypeAccept                 //Accept产生的
)

//func GetTimestamp() int64 {
//	return time.Now().UnixNano() / 1000_000_000
//}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if handler == nil {
			} else {
				handler(err)
			}
		}
	}()
	fun()
}

func Go(fu func()) {
	go Try(fu, nil)
}
