package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据......")

	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err:", err)
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//	反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
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

	//发送data
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail:", err)
		return
	}

	return
}
