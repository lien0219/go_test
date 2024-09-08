package process

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/client/utils"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
	"net"
	"os"
)

// 登录成功后
func ShowMenu() {
	fmt.Println("------恭喜xxx登录成功------")
	fmt.Println("------1.显示在线用户列表------")
	fmt.Println("------2.发送消息------")
	fmt.Println("------3.信息列表------")
	fmt.Println("------4.退出系统------")
	fmt.Println("请选择（1-4）：")

	var key int
	var content string

	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		//fmt.Println("发送消息")
		fmt.Println("你想说什么？：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选择项不对。。。")
	}
}

// 和服务端保持通讯
func ServerProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err:", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			UpdateUserStatus(&notifyUserStatusMes)
			//通知有人上线
		default:
			fmt.Println("服务端返回了未知的消息类型")
		}
		fmt.Println("mes:", mes)
	}
}
