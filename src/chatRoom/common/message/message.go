package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

// 用户状态
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息时间
}
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
type LoginResMes struct {
	Code    int    `json:"code"` //状态码：500未注册 200成功
	UserIds []int  //保护用户id的切片
	Error   string `json:"error"` //错误信息
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// 配合服务器推送客户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
