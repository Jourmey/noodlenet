package noodle

import (
	"net/http"
	"noodlenet/deps/src/github.com/gorilla/websocket"
	"noodlenet/deps/src/github.com/zapLogger"
	"sync/atomic"
	"time"
)

//type IConnect interface {
//	read()
//	write()
//	Stop()
//}

type WsConnect struct {
	id      uint32 //唯一标示
	addr    string
	url     string
	timeout int //传输超时

	connTyp  ConnType //通道类型
	conn     *websocket.Conn
	upgrader *websocket.Upgrader
	listener *http.Server

	handler IMsgHandler
	cWrite  chan *Message //写入通道

	init      bool
	available bool
	stop      int32 //停止标记
	lastTick  time.Time
}

func (w *WsConnect) write() {
	var m *Message
	tick := time.NewTimer(time.Second * time.Duration(w.timeout))
	//gm := GMManager.GetGMessage()
	for true {
		if !w.IsStop() || m == nil {
			select {
			case m = <-w.cWrite:
			//case <-gm.c: //广播
			//	if gm.fun == nil || gm.fun(w) {
			//		m = gm.msg
			//	}
			//	gm = GMManager.GetGMessage()
			case <-tick.C:
				if w.isTimeout() {
					w.Stop()
				}
			}
		}
		if m == nil || m.Data == nil {
			continue
		}
		err := w.conn.WriteMessage(websocket.BinaryMessage, m.Data)
		if err != nil {
			zapLogger.Errorf("[%v] write err:%v", w.id, err)
			break
		} else {
			zapLogger.Infof("[%v] write success:%+v", w.id, m)
		}
		m = nil
		w.lastTick = time.Now()
	}
}

func (w *WsConnect) Send(m *Message) (re bool) {
	if m == nil || !w.available {
		return
	}
	Try(func() {
		w.cWrite <- m
		re = true
	}, func(i interface{}) {
		zapLogger.Errorf("Send Message Failed ,err = ", i)
		re = false
	})
	return
}

func (w *WsConnect) Stop() {
	if atomic.CompareAndSwapInt32(&w.stop, 0, 1) {
		Go(func() {
			if w.init {
				w.handler.OnDelMsgQue(w)
			}
			w.available = false
			if w.cWrite != nil {
				close(w.cWrite)
			}

			ConnectMapSync.Lock()
			delete(ConnectMap, w.id)
			ConnectMapSync.Unlock()
			zapLogger.Infof("[%v] close ", w.id)
		})
	}
}
func (w *WsConnect) IsStop() bool {
	return w.stop == 1
}

func (w *WsConnect) read() {
	for !w.IsStop() {
		_, data, err := w.conn.ReadMessage()
		if err != nil {
			break
		} else {
			zapLogger.Infof("[%v] read success:%+v", w.id, data)
		}
		if !w.processMsg(w, &Message{Data: data}) {
			break
		}
		w.lastTick = time.Now()
	}
}

func (w *WsConnect) processMsg(msgque *WsConnect, msg *Message) bool {
	f := w.handler.GetHandlerFunc(msgque, msg)
	if f == nil {
		f = w.handler.OnProcessMsg
	}
	return f(msgque, msg)
}

func NewConnect(addr, url string, handler IMsgHandler) *WsConnect {
	con := new(WsConnect)
	con.id = atomic.AddUint32(&ConnectId, 1)
	con.addr = addr
	con.url = url
	con.connTyp = ConnTypeListen
	con.handler = handler
	con.listener = &http.Server{Addr: addr}

	ConnectMapSync.Lock()
	ConnectMap[con.id] = con
	ConnectMapSync.Unlock()
	zapLogger.Infof("[%v] New Connect addr:%s url:%s", con.id, addr, url)
	return con
}

func (w *WsConnect) ListenAndServe() {
	w.upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc(w.url, func(hw http.ResponseWriter, hr *http.Request) {
		c, err := w.upgrader.Upgrade(hw, hr, nil)
		if err != nil {
			zapLogger.Errorf("[%v] accept failed err:%v", w.id, err)
		} else {
			Go(func() {
				msgque := newWsAccept(c, w.handler)
				if w.handler.OnNewMsgQue(msgque) {
					msgque.init = true
					msgque.available = true
					Go(func() {
						zapLogger.Infof("[%v] process read", msgque.id)
						msgque.read()
						zapLogger.Infof("[%v] process read end", msgque.id)
					})
					Go(func() {
						zapLogger.Infof("[%v] process write", msgque.id)
						msgque.write()
						zapLogger.Infof("[%v] process write end", msgque.id)
					})
				} else {
					msgque.Stop()
				}
			})
		}
	})
	w.listener.ListenAndServe()
}

func (w *WsConnect) isTimeout() bool {
	d := time.Now().Sub(w.lastTick)
	return int(d.Seconds()) > w.timeout
}

func newWsAccept(conn *websocket.Conn, handler IMsgHandler) *WsConnect {
	con := new(WsConnect)
	con.id = atomic.AddUint32(&ConnectId, 1)
	con.timeout = 60
	con.cWrite = make(chan *Message, 64)
	con.connTyp = ConnTypeAccept
	con.handler = handler
	con.conn = conn

	ConnectMapSync.Lock()
	ConnectMap[con.id] = con
	ConnectMapSync.Unlock()
	zapLogger.Infof("[%v] new WsAccept from addr:%s", con.id, conn.RemoteAddr().String())
	return con
}
