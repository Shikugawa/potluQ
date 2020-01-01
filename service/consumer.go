package service

import (
	"github.com/Shikugawa/potluq/external"
	"github.com/Shikugawa/potluq/message"
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
