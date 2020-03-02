package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"data"` // 消息的内容
}

// 定义两个消息类型，后续在增加
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户Id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名称
}

type LoginResMes struct {
	Code    int    `json:"code"` // 返回状态码 500 表示用户未组册，200表示登录成功
	UserIds []int  // 增加字段，保存用户id的切片
	Error   string `json:"error"` // 返回的错误信息
}

// 用户注册包协议结构
type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示用户被注册，200表示注册成功
	Error string `json:"error"` // 返回的错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 聊天消息
type SmsMes struct {
	Content string `json:"content"`
	User
}
