package process

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis/src/chatRoom/common/message"
	"github.com/gomodule/redigo/redis/src/chatRoom/server/model"
	"github.com/gomodule/redigo/redis/src/chatRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//	区别用户的Conn标识
	UserId int
}

// 通知所有在线的用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// notifyUserStatusMes 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	mes.Data = string(data)

	//对mes再次序列化，发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	//Transform实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err:", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册错误"
		}
	} else {
		registerResMes.Code = 200
	}
	return
}

// 处理登录
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	//redis校验
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		//用户不存在
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		//将登陆成功的用户的userid给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他在线用户某某上线
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//当前在线用户的id放入loginResMes.UserIds
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

		fmt.Println(user, "登录成功")
	}

	//如果用户id=100，密码=123456则合法，否则不合法
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200
	//} else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在，请注册再使用"
	//}

	//loginResMess序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err:", err)
		return
	}

	resMes.Data = string(data)

	//对resMes序列化并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
