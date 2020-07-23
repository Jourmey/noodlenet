package main

import "noodlenet/noodle"

type ChatInfoFunc func(connect *noodle.WsConnect, c *ChatInfo)

const (
	LogIn2S uint8 = iota // 登陆
	LogIn2C              // 登陆结果
	Chat2S               // 单数据
	BChat2S              // 广播数据
)

type UserID string

var BroadCastID UserID = "xxxx-xxxx"

type ChatInfo struct {
	ID   UserID
	Info string
}

func NewChat(id, info string) *ChatInfo {
	c := new(ChatInfo)
	c.ID = UserID(id)
	c.Info = info
	return c
}

func NewChatInfo(msg *noodle.Message) *ChatInfo {
	c := new(ChatInfo)
	c.ID = UserID(msg.Data[:16])
	c.Info = string(msg.Data[16:])

	return c
}
