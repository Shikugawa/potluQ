package message

type Media int

const (
	YouTube Media = iota
)

type QueueMessage struct {
	UserName  string
	ClubName  string
	MediaType Media
	Url       string
}
