package infra

import (
	"github.com/go-redis/redis"
)

type RedisHandler struct {
	Conn *redis.Client
}

func InitRedisHandler(host, port string) *RedisHandler {
	return &RedisHandler{
		Conn: redis.NewClient(&redis.Options{
			Addr: host + ":" + port,
		}),
	}
}
