package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
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
	Code  int    `json:"code"`  //状态码：500未注册 200成功
	Error string `json:"error"` //错误信息
}
type RegisterMes struct {
}
