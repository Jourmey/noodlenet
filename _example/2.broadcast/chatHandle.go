package main

import (
	log "noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle"
	"noodlenet/noodle/msg"
)

type ChatInfoFunc func(connect *noodle.WsConnect, c *msg.Pb)

const (
	LogIn2S uint32 = iota // 登陆
	LogIn2C               // 登陆结果
	Chat2S                // 单数据
	BChat2S               // 广播数据
)

type UserID string

var BroadCastID UserID = "xxxx-xxxx"

type ChatHandler struct {
	msgMap map[uint32]ChatInfoFunc
}

func NewChatHandler() *ChatHandler {
	c := new(ChatHandler)
	c.msgMap = make(map[uint32]ChatInfoFunc, 0)
	return c
}

func (d *ChatHandler) OnNewConnect(c *noodle.WsConnect) bool {
	log.Infof("[%d] New连接", c.ID)
	return true
}

func (d *ChatHandler) OnDelConnect(c *noodle.WsConnect) {
	log.Infof("[%d] Del连接", c.ID)
}

func (d *ChatHandler) HandlerFunc(c *noodle.WsConnect, msg *msg.Pb) {
	if msg == nil {
		log.Infof("[%d] msg == nil", c.ID)
		return
	}

	if fu, ok := d.msgMap[msg.Cmd]; ok {
		noodle.Try(func() {
			fu(c, msg)
		}, func(i interface{}) {
			log.Error("[%d] HandlerFunc err = %v", c.ID, i)
		})
	}
}

func (d *ChatHandler) Register(cmd uint32, fun ChatInfoFunc) {
	d.msgMap[cmd] = fun
}
