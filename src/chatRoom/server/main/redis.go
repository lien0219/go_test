package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, MaxIdle int, MaxActive int, IdleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     MaxIdle,     //最大空连接数量
		MaxActive:   MaxActive,   //和数据库最大的连接数  0是不限制
		IdleTimeout: IdleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
