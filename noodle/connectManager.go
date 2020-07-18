package noodle

import (
	"sync"
	"sync/atomic"
)

type connectManager struct {
	connectId      uint32 //消息队列id
	connectMapSync sync.Mutex
	connectMap     map[uint32]*WsConnect
}

func (m *connectManager) init() {
	m.connectMap = make(map[uint32]*WsConnect)
}

func (m *connectManager) getConnectID() uint32 {
	return atomic.AddUint32(&m.connectId, 1)
}

func (m *connectManager) addConnect(con *WsConnect) {
	m.connectMapSync.Lock()
	m.connectMap[con.ID] = con
	m.connectMapSync.Unlock()
}

func (m *connectManager) deleteConnect(w *WsConnect) {
	m.connectMapSync.Lock()
	delete(m.connectMap, w.ID)
	m.connectMapSync.Unlock()
}
