package db

import "github.com/gomodule/redigo/redis"

var RedisPoolInstance *redis.Pool

func init() {
	RedisPoolInstance = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		DialContext:     nil,
		TestOnBorrow:    nil,
		MaxIdle:         8,   // 最大空闲数
		MaxActive:       10,  // 最大连接数 -- 0表示无限制
		IdleTimeout:     100, // 最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
