package role

type Role int

const (
	DefaultRole   Role = iota // can access to list of Clubs, Default role
	JukeBoxRole               // can manage all operation about music
	RequesterRole             // can enqueue to music queue
)
