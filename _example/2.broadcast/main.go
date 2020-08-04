package main

import (
	"noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle"
	"noodlenet/noodle/msg"
	"sync"
)

var gamerMap = map[uint32]*noodle.WsConnect{}
var currentID uint32 = 0
var lock sync.Mutex

func main() {
	//zapLogger.InitLogger("debug", zapLogger.Logc)

	hand := NewChatHandler()
	hand.Register(LogIn2S, HandleLogIn2S)
	hand.Register(Chat2S, HandleChat2S)
	hand.Register(BChat2S, HandleBChat2S)
	noodle.NewConnect(":5000", "/", hand).SetTimeout(30).ListenAndServe()
}

func HandleLogIn2S(connect *noodle.WsConnect, c *msg.Pb) { //登陆
	zapLogger.Infof("登陆信息 , data = %s", c.Data)
	s2c := &msg.Pb{
		Cmd:  LogIn2C,
		Data: []byte("登陆成功！"),
	}
	lock.Lock()
	currentID++
	s2c.Index = currentID
	lock.Unlock()

	if connect.Send(s2c) {
		gamerMap[s2c.Index] = connect
	}
}

func HandleBChat2S(connect *noodle.WsConnect, c *msg.Pb) { //广播
	zapLogger.Infof("广播信息 , data = %s", c.Data)

	s2c := &msg.Pb{
		Cmd:   BChat2C,
		Data:  c.Data,
		Index: c.Index,
	}

	noodle.GMManager.Send(s2c, func(con *noodle.WsConnect) bool {
		//return connect != con
		return true
	})
}

func HandleChat2S(connect *noodle.WsConnect, c *msg.Pb) { //单播
	zapLogger.Infof("单播信息 , data = %s", c.Data)

	if mq, ok := gamerMap[c.Index]; ok && mq.Available() {
		s2c := &msg.Pb{
			Cmd:  Chat2C,
			Data: c.Data,
		}
		s2cR := &msg.Pb{
			Cmd: Chat2C,
		}
		if mq.Send(s2c) {
			s2cR.Data = []byte("发送成功")
		} else {
			s2cR.Data = []byte("发送失败")
		}
		connect.Send(s2cR)

	} else {
		s2c := &msg.Pb{
			Cmd:  Chat2C,
			Data: []byte("用户已离线"),
		}
		connect.Send(s2c)
	}
}
