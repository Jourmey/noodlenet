package main

import (
	log "noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle"
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

func (d *EchoHandler) MessageFromBytes(data []byte) (*noodle.Message, error) {
	return &noodle.Message{Data: data}, nil
}

func (d *EchoHandler) MessageToBytes(msg *noodle.Message) ([]byte, error) {
	return msg.Data, nil
}

func (d *EchoHandler) HandlerFunc(c *noodle.WsConnect, msg *noodle.Message) {
	_ = c.Send(msg)
}
