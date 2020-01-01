package infra

import (
	"log"

	"github.com/Shikugawa/potluq/external"
	"github.com/Shikugawa/potluq/message"
)

type Worker struct {
	redisClient external.RedisHandler
	queue       *chan message.QueueMessage
	retryTimes  int
}

func InitWorker(redisClient external.RedisHandler, queue *chan message.QueueMessage, retry int) *Worker {
	return &Worker{
		redisClient: redisClient,
		queue:       queue,
		retryTimes:  retry,
	}
}

func (worker *Worker) Start() {
	for {
		select {
		case message := <-*worker.queue:
			err := worker.Enqueue(&message)
			// try to enqueue retryTimes
			if err != nil {
				times := 0
				for times <= worker.retryTimes {
					if times == worker.retryTimes {
						log.Fatalln("failed to enqueue music")
						break
					}
					err = worker.Enqueue(&message)
					if err == nil {
						break
					}
					times++
				}
			}
		default:
		}
	}
}

func (worker *Worker) Enqueue(message *message.QueueMessage) error {
	if err := worker.redisClient.EnqueueMusic(message); err != nil {
		return err
	}
	return nil
}
