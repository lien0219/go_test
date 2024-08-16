package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
	"net"
)

func Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId = %d userPwd = %s\n", userId, userPwd)
	//return nil

	//	连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	defer conn.Close()

	//	通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	mes.Data = string(data)

	//	mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	//	data长度发送给服务器
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail:", err)
		return
	}

	fmt.Println("客户端发送消息的长度ok")

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail:", err)
		return
	}

	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20...")
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err:", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
