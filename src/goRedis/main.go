package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空连接数量
		MaxActive:   0,   //和数据库最大的连接数  0是不限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//	取出pool一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "李四")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//	取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("r=", r)
}
