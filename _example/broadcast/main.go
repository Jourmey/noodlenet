package main

import (
	"errors"
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
	if data != nil {
		return &noodle.Message{Data: data}, nil
	}
	return nil, errors.New("data is nil")
}

func (d *EchoHandler) MessageToBytes(msg *noodle.Message) ([]byte, error) {
	if msg != nil && msg.Data != nil {
		return msg.Data, nil
	}
	return nil, errors.New("msg is nil")
}

func (d *EchoHandler) HandlerFunc(c *noodle.WsConnect, msg *noodle.Message) {
	if string(msg.Data) == "close" {
		c.Stop()
	} else if string(msg.Data) == "broadcast" {
		noodle.GMManager.Send(msg, func(connect *noodle.WsConnect) bool {
			log.Infof("[%d] 发送广播数据  message = %+v", connect.ID, msg)
			return true
		})
	} else {
		log.Infof("[%d] 接收数据  message = %+v", c.ID, msg)
		_ = c.Send(msg)
		log.Infof("[%d] 发送数据  message = %+v", c.ID, msg)
	}
}
