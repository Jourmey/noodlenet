package noodle

import (
	"noodlenet/noodle/msg"
	"sync"
)

type gMessage struct {
	c       chan struct{} // 触发广播使用
	message *msg.Pb
	fun     func(connect *WsConnect) bool
}

const GMessageCap uint16 = 65535

type gMessageManager struct {
	gMessageArray [GMessageCap]*gMessage
	gId           uint16
	gSync         sync.Mutex
}

func (g *gMessageManager) init() {
	g.gId = 0
	g.gMessageArray[g.gId] = &gMessage{c: make(chan struct{})}
}

func (g *gMessageManager) Send(msg *msg.Pb, fun func(connect *WsConnect) bool) {
	if msg == nil {
		return
	}
	c := make(chan struct{})
	g.gSync.Lock()
	gm := g.gMessageArray[g.gId]
	g.gId = (g.gId + 1) % GMessageCap
	g.gMessageArray[g.gId] = &gMessage{c: c}
	g.gSync.Unlock()
	gm.message = msg
	gm.fun = fun
	close(gm.c)
}

func (g *gMessageManager) getGMessage() *gMessage {
	return g.gMessageArray[g.gId]
}
