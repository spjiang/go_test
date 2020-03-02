package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 定义一个全局的pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address)
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}
}
