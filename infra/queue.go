package infra

import (
	"github.com/Shikugawa/potraq/ent"
)

type Media int

const (
	YouTube Media = iota
)

type QueueMessage struct {
	User      *ent.User
	MediaType Media
	Url       string
}

type Queue struct {
	messages []QueueMessage
}
