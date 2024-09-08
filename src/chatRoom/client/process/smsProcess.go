package process

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/client/utils"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
)

type SmsProcess struct {
}

// 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal failed err:", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal failed err:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err:", err)
		return
	}

	return
}
