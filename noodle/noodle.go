package noodle

var ConnectManager connectManager //连接管理
var GMManager gMessageManager     //广播管理

type ConnType uint8

const (
	ConnTypeListen ConnType = iota //监听
	ConnTypeConn                   //连接产生的
	ConnTypeAccept                 //Accept产生的
)

func init() {
	ConnectManager.init()
	GMManager.init()
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if handler == nil {
			} else {
				handler(err)
			}
		}
	}()
	fun()
}

func Go(fu func()) {
	go Try(fu, nil)
}
