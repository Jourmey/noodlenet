package noodle

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"noodlenet/deps/src/github.com/gorilla/websocket"
	log "noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/noodle/msg"
	"sync/atomic"
	"time"
)

//type IConnect interface {
//	read()
//	write()
//	Stop()
//}

type WsConnect struct {
	ID      uint32 //唯一标示
	addr    string
	url     string
	timeout int //传输超时

	connTyp  ConnType //通道类型
	conn     *websocket.Conn
	upgrader *websocket.Upgrader
	listener *http.Server

	handler IMsgHandler
	cWrite  chan *msg.Pb //写入通道

	init      bool
	available bool
	stop      int32 //停止标记
	lastTick  time.Time

	UserToken interface{}
}

func NewConnect(addr, url string, handler IMsgHandler) *WsConnect {
	con := new(WsConnect)
	con.ID = ConnectManager.getConnectID()
	con.addr = addr
	con.url = url
	con.connTyp = ConnTypeListen
	con.handler = handler
	con.listener = &http.Server{Addr: addr}

	ConnectManager.addConnect(con)

	log.Debugf("[%v] New Connect addr:%s url:%s", con.ID, addr, url)
	return con
}

func (w *WsConnect) Available() bool {
	return w.available
}

func (w *WsConnect) SetTimeout(timeout int) *WsConnect {
	w.timeout = timeout
	return w
}

func (w *WsConnect) Send(m *msg.Pb) (re bool) {
	if m == nil || !w.available {
		return
	}
	Try(func() {
		w.cWrite <- m
		re = true
	}, func(i interface{}) {
		log.Errorf("Send Message Failed ,err = ", i)
		re = false
	})
	return
}

func (w *WsConnect) Stop() {
	if atomic.CompareAndSwapInt32(&w.stop, 0, 1) {
		if w.init {
			w.handler.OnDelConnect(w)
		}
		w.available = false
		if w.cWrite != nil {
			close(w.cWrite)
		}

		ConnectManager.deleteConnect(w)

		log.Debugf("[%v] close ", w.ID)
	}
}

func (w *WsConnect) IsStop() bool {
	return w.stop == 1
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
			log.Errorf("[%v] accept failed err:%v", w.ID, err)
		} else {
			Go(func() {
				connect := newWsAccept(c, w.handler, w.timeout)
				if w.handler.OnNewConnect(connect) {
					connect.init = true
					connect.available = true
					Go(func() {
						log.Debugf("[%v] process read", connect.ID)
						connect.read()
						log.Debugf("[%v] process read end", connect.ID)
					})
					Go(func() {
						log.Debugf("[%v] process write", connect.ID)
						connect.write()
						log.Debugf("[%v] process write end", connect.ID)
					})
				} else {
					connect.Stop()
				}
			})
		}
	})
	w.listener.ListenAndServe()
}

func (w *WsConnect) read() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("[%v] read panic err:%v", w.ID, err)
		}
		w.Stop()
	}()

	for !w.IsStop() {
		_, data, err := w.conn.ReadMessage()
		if err != nil {
			log.Debugf("[%v] read failed:%+v", w.ID, err.Error())
			break
		} else {
			log.Debugf("[%v] read success:%+v", w.ID, data)
		}

		var m msg.Pb
		err = proto.Unmarshal(data, &m)
		if err != nil {
			continue
		}
		w.handler.HandlerFunc(w, &m)
		w.lastTick = time.Now()
	}
}

func (w *WsConnect) write() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("[%v] write panic err:%v", w.ID, err)
		}
		if w.conn != nil {
			w.conn.Close()
		}
		w.Stop()
	}()

	var m *msg.Pb
	gm := GMManager.getGMessage()
	tick := time.NewTimer(time.Second * time.Duration(w.timeout))
	for !w.IsStop() {
		select {
		case m = <-w.cWrite:
		case <-gm.c: //广播
			if gm.fun == nil || gm.fun(w) {
				m = gm.message
			}
			gm = GMManager.getGMessage()
		case <-tick.C:
			if isTimeout, offTimes := w.isTimeout(); isTimeout {
				w.Stop()
			} else {
				tick.Reset(time.Second * time.Duration(offTimes))
			}
		}

		data, err := proto.Marshal(m)
		if err != nil {
			log.Debugf("[%v] Marshal failed:%+v", w.ID, err)
			continue
		}
		err = w.conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Errorf("[%v] write err:%v", w.ID, err)
			break
		} else {
			log.Debugf("[%v] write success:%+v", w.ID, m)
		}
		m = nil
		w.lastTick = time.Now()
	}
}

func (w *WsConnect) isTimeout() (bool, int) {
	if w.timeout == 0 {
		return false, -1
	}

	d := time.Now().Sub(w.lastTick)
	p := int(d.Seconds())
	if p < w.timeout {
		return false, w.timeout - p
	}
	log.Debugf("[%v] timeout timeout:%v", w.ID, w.timeout)
	return true, p - w.timeout
}

func newWsAccept(conn *websocket.Conn, handler IMsgHandler, timeout int) *WsConnect {
	con := new(WsConnect)
	con.ID = ConnectManager.getConnectID()
	con.cWrite = make(chan *msg.Pb, 64)
	con.timeout = timeout
	con.connTyp = ConnTypeAccept
	con.handler = handler
	con.conn = conn

	ConnectManager.addConnect(con)

	log.Debugf("[%v] new WsAccept from addr:%s", con.ID, conn.RemoteAddr().String())
	return con
}
