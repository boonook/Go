package myredis

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func Conn() redis.Conn {
	////新建连接
	return pool.Get()
}
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     80,
		MaxActive:   12000, // max number of connections
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				fmt.Println("Connect to redis error", err)
				panic(err)
			} else {
				beego.Info("连接redis成功")
			}
			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func init() {
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")

	pool = newPool(server, password)
}
