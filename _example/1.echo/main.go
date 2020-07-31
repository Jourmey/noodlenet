package main

import (
	log "noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle"
	"noodlenet/noodle/msg"
)

func main() {
	hand := new(EchoHandler)
	noodle.NewConnect(":5000", "/", hand).SetTimeout(30).ListenAndServe()
}

type EchoHandler struct {
}

func (d *EchoHandler) OnNewConnect(c *noodle.WsConnect) bool {
	log.Infof("[%d] New连接", c.ID)
	return true
}

func (d *EchoHandler) OnDelConnect(c *noodle.WsConnect) {
	log.Infof("[%d] Del连接", c.ID)
}

func (d *EchoHandler) HandlerFunc(c *noodle.WsConnect, msg *msg.Pb) {
	log.Infof("[%d] 接收数据  message = %+v", c.ID, msg)
	_ = c.Send(msg)
	log.Infof("[%d] 发送数据  message = %+v", c.ID, msg)
}
