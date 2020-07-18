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

func (d *EchoHandler) OnProcessMsg(c *noodle.WsConnect, msg *noodle.Message) bool {
	log.Info("implement me")
	return true
}

func (d *EchoHandler) GetHandlerFunc(c *noodle.WsConnect, msg *noodle.Message) noodle.HandlerFunc {
	return func(c *noodle.WsConnect, msg *noodle.Message) bool {
		if string(msg.Data) == "close" {
			c.Stop()
			return true
		} else if string(msg.Data) == "broadcast" {
			noodle.GMManager.Send(msg, func(connect *noodle.WsConnect) bool {
				log.Infof("[%d] 发送广播数据  message = %+v", connect.ID, msg)
				return true
			})

			return true
		} else {
			log.Infof("[%d] 接收数据  message = %+v", c.ID, msg)
			b := c.Send(msg)
			log.Infof("[%d] 发送数据  message = %+v", c.ID, msg)
			return b
		}
	}
}
