package connet

import (
	"net/http"
	"noodlenet/deps/src/github.com/gorilla/websocket"
	"noodlenet/deps/src/github.com/zapLogger"
	"noodlenet/message"
)

type IConnect interface {
	read()
	write()
	Stop()
}

type wsConnect struct {
	id   uint32 //唯一标示
	addr string
	url  string

	conn     *websocket.Conn
	upgrader *websocket.Upgrader
	listener *http.Server

	handler IMsgHandler

	cwrite chan *message.Message //写入通道

	init      bool
	available bool
}

func (w *wsConnect) write() {
	var m *message.Message
	for true {
		if m == nil {
			select {
			case m = <-w.cwrite:
			}
			if m == nil || m.Data == nil {
				continue
			}
			err := w.conn.WriteMessage(websocket.BinaryMessage, m.Data)
			if err != nil {
				zapLogger.Errorf("msgque write id:%v err:%v", w.id, err)
				break
			}
			m = nil
		}
	}
}

func (w *wsConnect) Stop() {
}

func (w *wsConnect) read() {
	for {
		_, data, err := w.conn.ReadMessage()
		if err != nil {
			break
		}
		if !w.processMsg(w, &message.Message{Data: data}) {
			break
		}
	}
}

func (w *wsConnect) processMsg(msgque IConnect, msg *message.Message) bool {
	f := w.handler.GetHandlerFunc(msgque, msg)
	if f == nil {
		f = w.handler.OnProcessMsg
	}
	return f(msgque, msg)
}

func NewWsAccept(addr, url string, handler IMsgHandler) IConnect {
	con := new(wsConnect)
	con.addr = addr
	con.url = url
	con.handler = handler

	con.listener = &http.Server{Addr: addr}

	con.upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc(con.url, func(hw http.ResponseWriter, hr *http.Request) {
		c, err := con.upgrader.Upgrade(hw, hr, nil)
		if err != nil {
			zapLogger.Errorf("accept failed msgque:%v err:%v", con.id, err)
		} else {
			Go(func() {
				msgque := newWsAccept(c, con.handler)
				if con.handler.OnNewMsgQue(msgque) {
					con.init = true
					con.available = true
					Go(func() {
						msgque.read()
					})
					Go(func() {
						msgque.write()
					})
				} else {
					msgque.Stop()
				}
			})
		}
	})
	con.listener.ListenAndServe()

	return con
}

func newWsAccept(conn *websocket.Conn, handler IMsgHandler) IConnect {
	msgque := new(wsConnect)
	msgque.id = 1
	msgque.handler = handler
	msgque.conn = conn
	msgque.cwrite = make(chan *message.Message, 64)

	return msgque
}

type HandlerFunc func(msgque IConnect, msg *message.Message) bool
type IMsgHandler interface {
	OnNewMsgQue(msgque IConnect) bool                                 //新的消息队列
	OnDelMsgQue(msgque IConnect)                                      //消息队列关闭
	OnProcessMsg(msgque IConnect, msg *message.Message) bool          //默认的消息处理函数
	//OnConnectComplete(msgque IConnect, ok bool) bool                  //连接成功
	GetHandlerFunc(msgque IConnect, msg *message.Message) HandlerFunc //根据消息获得处理函数
}

func Go(fu func()) {
	go fu()
}
