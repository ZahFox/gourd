package command 

// Action is what the gourd client is requesting to be done
type Action int

const (
	// NOACTION is a request with no action
	NOACTION Action = iota
	// PING is a request to recieve a PONG response
	PING
	// ECHO is a request to have text echoed back
	ECHO
)