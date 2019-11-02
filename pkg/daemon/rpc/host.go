package rpc

// Host accepts and responds to rpc commands for the local gourdd
type Host struct {
}

// Ping responds with pong
func (c *Host) Ping(_ interface{}, reply *string) error {
	*reply = "PONG"
	return nil
}

// Echo responds with the message that was sent to it
func (c *Host) Echo(message *string, reply *string) error {
	*reply = *message
	return nil
}
