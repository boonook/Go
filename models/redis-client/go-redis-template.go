package redis_client

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

server := beego.AppConfig.String("cache::server")
password := beego.AppConfig.String("cache::password")

func GetClient() (clint *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: password // no password set
	})

	pong, err := client.Ping().Result()
	log.Infof(pong, err)
	return client
}
