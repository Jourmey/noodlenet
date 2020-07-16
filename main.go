package main

import (
	"noodlenet/connet"
	"noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/message"
	"time"
)

func main() {
	hand := new(Dandler)

	_ = connet.NewWsAccept(":5000", ":5000", hand)

	time.Sleep(1000 * time.Second)

	//ic.Stop()
}

type Dandler struct {
}

func (d *Dandler) OnNewMsgQue(msgque connet.IConnect) bool {
	zapLogger.Info("implement me")
	return true
}

func (d *Dandler) OnDelMsgQue(msgque connet.IConnect) {
	zapLogger.Info("implement me")

}

func (d *Dandler) OnProcessMsg(msgque connet.IConnect, msg *message.Message) bool {
	zapLogger.Info("implement me")
	return true
}

func (d *Dandler) GetHandlerFunc(msgque connet.IConnect, msg *message.Message) connet.HandlerFunc {
	return d.OnProcessMsg
}
