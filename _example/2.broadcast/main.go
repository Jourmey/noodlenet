package main

import (
	"noodlenet/noodle"
)

func main() {
	hand := NewChatHandler()
	hand.Register(LogIn2S, func(connect *noodle.WsConnect, c *ChatInfo) {

	})
	hand.Register(Chat2S, func(connect *noodle.WsConnect, c *ChatInfo) {

	})
	hand.Register(BChat2S, func(connect *noodle.WsConnect, c *ChatInfo) {

	})
	noodle.NewConnect(":5000", "/", hand).SetTimeout(30).ListenAndServe()
}
