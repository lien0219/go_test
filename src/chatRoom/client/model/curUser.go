package model

import (
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
