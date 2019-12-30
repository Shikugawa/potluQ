package infra

import (
	"encoding/json"

	"github.com/Shikugawa/potraq/external"
	"github.com/Shikugawa/potraq/message"
	"github.com/go-redis/redis"
)

type RedisHandler struct {
	Conn *redis.Client
}

func InitRedisHandler(host, port string) external.RedisHandler {
	return &RedisHandler{
		Conn: redis.NewClient(&redis.Options{
			Addr: host + ":" + port,
		}),
	}
}

func (handler *RedisHandler) EnqueueMusic(message *message.QueueMessage) error {
	if err := handler.Conn.LPush(message.ClubName, *message).Err(); err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) DequeueMusic(clubName string) (*message.QueueMessage, error) {
	var message message.QueueMessage
	str, err := handler.Conn.RPop(clubName).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(str), message); err != nil {
		return nil, err
	}
	return &message, nil
}
