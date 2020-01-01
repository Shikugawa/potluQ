package external

import "github.com/Shikugawa/potluq/message"

type RedisHandler interface {
	EnqueueMusic(message *message.QueueMessage) error
	DequeueMusic(clubName string) (*message.QueueMessage, error)
}
