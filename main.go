package main

import (
	"noodlenet/noodle"
	"noodlenet/deps/src/github.com/zapLogger"
)

func main() {
	hand := new(Dandler)
	wc := noodle.NewConnect(":5000", "/", hand)
	wc.ListenAndServe()
}

type Dandler struct {
}

func (d *Dandler) OnNewMsgQue(c *noodle.WsConnect) bool {
	zapLogger.Info("implement me")
	return true
}

func (d *Dandler) OnDelMsgQue(c *noodle.WsConnect) {
	zapLogger.Info("implement me")
}

func (d *Dandler) OnProcessMsg(c *noodle.WsConnect, msg *noodle.Message) bool {
	zapLogger.Info("implement me")
	return true
}

func (d *Dandler) GetHandlerFunc(c *noodle.WsConnect, msg *noodle.Message) noodle.HandlerFunc {
	return func(c *noodle.WsConnect, msg *noodle.Message) bool {
		if string(msg.Data) == "close" {
			c.Stop()
			return true
		} else {
			return c.Send(msg)
		}
	}
}
