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

func (d *EchoHandler) ParserMsg(data []byte) (*noodle.Message, error) {
	//if len(data) < noodle.MessageHeadSize {
	//	return nil, noodle.ErrMsgLenTooShort
	//}
	//m := (*noodle.Message)(unsafe.Pointer(&data[0]))
	//if m.Len > noodle.MaxMsgDataSize {
	//	return nil, noodle.ErrMsgLenTooLong
	//}
	return &noodle.Message{Data: data}, nil
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
