package process

import (
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/client/model"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
)

// 客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

// 客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 处理返回的NotifyUserStatusMes
func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
