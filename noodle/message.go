package noodle

const (
	MessageHeadSize        = 12
	MaxMsgDataSize  uint32 = 1024 * 1024
)

type Message struct {
	Index uint16 //序号
	Len   uint32 //数据长度
	Error uint16 //错误码
	Cmd   uint8  //命令
	Act   uint8  //动作
	Flags uint16 //标记
	Data  []byte //消息数据
}

var (
	ErrOk             = NewError("正确", 0)
	ErrDBErr          = NewError("数据库错误", 1)
	ErrProtoPack      = NewError("协议解析错误", 2)
	ErrProtoUnPack    = NewError("协议打包错误", 3)
	ErrMsgPackPack    = NewError("msgpack打包错误", 4)
	ErrMsgPackUnPack  = NewError("msgpack解析错误", 5)
	ErrPBPack         = NewError("pb打包错误", 6)
	ErrPBUnPack       = NewError("pb解析错误", 7)
	ErrJsonPack       = NewError("json打包错误", 8)
	ErrJsonUnPack     = NewError("json解析错误", 9)
	ErrCmdUnPack      = NewError("cmd解析错误", 10)
	ErrMsgLenTooLong  = NewError("数据过长", 11)
	ErrMsgLenTooShort = NewError("数据过短", 12)
	ErrHttpRequest    = NewError("http请求错误", 13)
	ErrConfigPath     = NewError("配置路径错误", 50)

	ErrFileRead   = NewError("文件读取错误", 100)
	ErrDBDataType = NewError("数据库数据类型错误", 101)
	ErrNetTimeout = NewError("网络超时", 200)

	ErrErrIdNotFound = NewError("错误没有对应的错误码", 255)
)

type Error struct {
	Id  uint16
	Str string
}

func (r *Error) Error() string {
	return r.Str
}

func NewError(s string, id uint16) *Error {
	err := &Error{id, s}
	return err
}
