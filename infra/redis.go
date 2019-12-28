package infra

import (
	"github.com/Shikugawa/potraq/interface/middleware"
	"github.com/go-redis/redis"
)

type RedisHandler struct {
	Conn *redis.Client
}

func InitRedisHandler(host, port string) middleware.RedisHandler {
	return &RedisHandler{
		Conn: redis.NewClient(&redis.Options{
			Addr: host + ":" + port,
		}),
	}
}
