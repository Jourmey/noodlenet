package main

import (
	"encoding/binary"
	"errors"
	log "noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle"
)

type ChatHandler struct {
	msgMap map[uint8]ChatInfoFunc
}

func NewChatHandler() *ChatHandler {
	c := new(ChatHandler)
	c.msgMap = make(map[uint8]ChatInfoFunc, 0)
	return c
}

func (d *ChatHandler) OnNewConnect(c *noodle.WsConnect) bool {
	log.Infof("[%d] New连接", c.ID)
	return true
}

func (d *ChatHandler) OnDelConnect(c *noodle.WsConnect) {
	log.Infof("[%d] Del连接", c.ID)
}

func (d *ChatHandler) MessageFromBytes(data []byte) (*noodle.Message, error) {
	if data == nil || len(data) < noodle.MessageHeadSize {
		return nil, errors.New("data is nil")
	}

	return &noodle.Message{
		Index: binary.LittleEndian.Uint16(data[0:]),
		Len:   binary.LittleEndian.Uint32(data[2:]),
		Error: binary.LittleEndian.Uint16(data[6:]),
		Cmd:   data[8],
		Act:   data[9],
		Flags: binary.LittleEndian.Uint16(data[10:]),
		Data:  data[noodle.MessageHeadSize:],
	}, nil
}

func (d *ChatHandler) MessageToBytes(msg *noodle.Message) ([]byte, error) {
	if msg == nil {
		return nil, errors.New("msg is nil")
	}
	n := 0
	if msg.Data != nil {
		n += len(msg.Data)
	}
	data := make([]byte, noodle.MessageHeadSize, noodle.MessageHeadSize+n)
	binary.LittleEndian.PutUint16(data[0:], msg.Index)
	binary.LittleEndian.PutUint32(data[2:], msg.Len)
	binary.LittleEndian.PutUint16(data[6:], msg.Error)
	data[8] = msg.Cmd
	data[9] = msg.Act
	binary.LittleEndian.PutUint16(data[10:], msg.Flags)

	if msg.Data != nil {
		data = append(data, msg.Data...)
	}
	return data, nil
}

func (d *ChatHandler) HandlerFunc(c *noodle.WsConnect, msg *noodle.Message) {
	if msg == nil {
		log.Infof("[%d] msg == nil", c.ID)
		return
	}
	if fu, ok := d.msgMap[msg.Cmd]; ok {
		noodle.Try(func() {
			fu(c, NewChatInfo(msg))
		}, func(i interface{}) {
			log.Error("[%d] HandlerFunc err = %v", c.ID, i)
		})
	}
}

func (d *ChatHandler) Register(cmd uint8, fun ChatInfoFunc) {
	d.msgMap[cmd] = fun
}
