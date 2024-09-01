package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/client/process"
	"os"
)

// 用户id & 用户密码
var (
	userId  int
	userPwd string
	key     int
	//loop    = true
	userName string
)

func main() {
	for true {
		fmt.Println("\t------欢迎登录聊天系统------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			//loop = false
			//	完成登录
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			//loop = false
			fmt.Println("请输用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的名字（nickname）")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	//if key == 1 {
	//
	//	//err := Login(userId, userPwd)
	//	//if err != nil {
	//	//	fmt.Println("登录失败")
	//	//} else {
	//	//	fmt.Println("登录成功")
	//	//}
	//} else if key == 2 {
	//	fmt.Println("用户进行注册")
	//}
}
