package main

import (
	"noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle"
	"noodlenet/noodle/msg"
)

var gamerMap = map[string]*noodle.WsConnect{}

func main() {
	zapLogger.InitLogger("debug", zapLogger.Logc)

	hand := NewChatHandler()
	hand.Register(LogIn2S, HandleLogIn2S)
	hand.Register(Chat2S, HandleChat2S)
	hand.Register(BChat2S, HandleBChat2S)
	noodle.NewConnect(":5000", "/", hand).SetTimeout(30).ListenAndServe()
}

func HandleLogIn2S(connect *noodle.WsConnect, c *msg.Pb) { //登陆
	gamerMap[string(c.Data)] = connect
	s2c := &msg.Pb{
		Cmd:  LogIn2C,
		Data: []byte("登陆成功"),
	}
	connect.Send(s2c)
}

func HandleChat2S(connect *noodle.WsConnect, c *msg.Pb) { //广播
	s2c := &msg.Pb{
		Cmd:  Chat2S,
		Data: []byte("广播:"),
	}
	s2c.Data = append(s2c.Data, c.Data...)

	noodle.GMManager.Send(s2c, func(con *noodle.WsConnect) bool {
		return connect != con
	})
}

func HandleBChat2S(connect *noodle.WsConnect, c *msg.Pb) { //单播
	if mq, ok := gamerMap[string(c.Data)]; ok && mq.Available() {
		mq.Send(c)
	} else {
		s2c := &msg.Pb{
			Cmd:  BChat2S,
			Data: []byte("用户已离线"),
		}
		connect.Send(s2c)
	}
}
