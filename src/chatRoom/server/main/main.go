package main

import (
	"fmt"
	"net"
)

//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	buf := make([]byte, 8096)
//	fmt.Println("读取客户端发送的数据......")
//
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//		fmt.Println("conn.Read err:", err)
//		//err = errors.New("read pkg header error")
//		return
//	}
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[0:4])
//
//	n, err := conn.Read(buf[:pkgLen])
//	if n != int(pkgLen) || err != nil {
//		//err = errors.New("read pkg body error")
//		return
//	}
//
//	//	反序列化
//	err = json.Unmarshal(buf[:pkgLen], &mes)
//	if err != nil {
//		fmt.Println("json.Unmarshal err:", err)
//		return
//	}
//	return
//}
//func writePkg(conn net.Conn, data []byte) (err error) {
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
//	//发送长度
//	n, err := conn.Write(buf[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(bytes) fail:", err)
//		return
//	}
//
//	//发送data
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write(data) fail:", err)
//		return
//	}
//
//	return
//}

//// 处理登录
//func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal fail err:", err)
//		return
//	}
//
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	var loginResMes message.LoginResMes
//
//	//如果用户id=100，密码=123456则合法，否则不合法
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		loginResMes.Code = 200
//	} else {
//		loginResMes.Code = 500
//		loginResMes.Error = "该用户不存在，请注册再使用"
//	}
//
//	//loginResMess序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail err:", err)
//		return
//	}
//
//	resMes.Data = string(data)
//
//	//对resMes序列化并发送
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail err:", err)
//		return
//	}
//	err = writePkg(conn, data)
//	return
//}

//// 处理客户端多种消息类型
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录
//		err = serverProcessLogin(conn, mes)
//	case message.RegisterMesType:
//	//处理注册
//	default:
//		fmt.Println("消息类型不存在......")
//	}
//
//	return
//}

// 处理客户端通讯
func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误 err:", err)
		return
	}
}
func main() {
	fmt.Println("服务器在8889端口监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	for {
		fmt.Println("等待客户端连接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			return
		}
		go process(conn)
	}
}
