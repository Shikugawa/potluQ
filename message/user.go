package message

type UserStatus int

const (
	JUKEBOX UserStatus = iota
	REQUESTER
)

type Credential struct {
	Status   UserStatus
	Email    string
	Password string
}
