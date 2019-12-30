package infra

import (
	"encoding/json"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/message"
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

func (handler *RedisHandler) EnqueueMusic(message *message.QueueMessage) error {
	if err := handler.Conn.LPush(message.ClubName, *message).Err(); err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) DequeueMusic(club *ent.Club) (*message.QueueMessage, error) {
	var message message.QueueMessage
	str, err := handler.Conn.RPop(club.Name).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(str), message); err != nil {
		return nil, err
	}
	return &message, nil
}
