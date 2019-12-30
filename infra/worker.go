package infra

import (
	"github.com/Shikugawa/potraq/message"
)

type Worker struct {
	redisHandler *RedisHandler
	queue        *chan message.QueueMessage
}

func InitWorker(redisHandler *RedisHandler, queue *chan message.QueueMessage) *Worker {
	return &Worker{
		redisHandler: redisHandler,
		queue:        queue,
	}
}

func (worker *Worker) Start() {
	for {
		select {
		case message := <-*worker.queue:
			worker.Enqueue(&message)
		default:
		}
	}
}

func (worker *Worker) Enqueue(message *message.QueueMessage) {
}
