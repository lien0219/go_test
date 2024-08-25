package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 根据用户id返回user实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := conn.Do("HGet", "users", id)
	if err != nil {
		//users哈希中没有找到对应的id
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	return
}

// 登录校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
