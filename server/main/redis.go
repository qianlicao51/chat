package main

import (
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(address string) {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address)
		},
		MaxIdle:     16,  //最大空闲连接
		MaxActive:   0,   //和数据库的最大连接数，0表示没限制
		IdleTimeout: 300, //最大空闲时间
	}
}
