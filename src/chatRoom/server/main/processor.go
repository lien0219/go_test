package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
	process2 "github.com/gomodule/redigo/redis/src/chatRoom/server/process"
	"github.com/gomodule/redigo/redis/src/chatRoom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 处理客户端多种消息类型
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//接收客户端群发消息
	fmt.Println("mes:", mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//完成转发群聊消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在......")
	}

	return
}

func (this *Processor) process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		//读取数据包
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端退出......")
				return
			} else {
				fmt.Println("readPkg err:", err)
				return
			}
		}
		//fmt.Println("mes:", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return
		}
	}
}
