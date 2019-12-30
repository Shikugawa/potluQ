package service

import (
	"github.com/Shikugawa/potraq/external"
	"github.com/Shikugawa/potraq/message"
)

type Consumer struct {
	Handler external.RedisHandler
}

func (consumer *Consumer) Consume(clubName string) (*message.QueueMessage, error) {
	music, err := consumer.Handler.DequeueMusic(clubName)
	if err != nil {
		return nil, err
	}
	return music, nil
}
