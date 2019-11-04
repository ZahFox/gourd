package rpc

import (
	"github.com/zahfox/gourd/pkg/command"
	"github.com/zahfox/gourd/pkg/utils"
)

// Host accepts and responds to rpc commands for the local gourdd
type Host struct {
}

// Ping responds with pong
func (c *Host) Ping(_ interface{}, reply *command.PingResponse) error {
	cmd := command.NewHostPing()
	utils.LogInfof("%s\n", cmd.String())
	*reply = command.PingResponse{
		ID:      cmd.ID,
		Error:   "",
		Message: "PONG",
	}
	return nil
}

// Echo responds with the message that was sent to it
func (c *Host) Echo(message *string, reply *command.EchoResponse) error {
	cmd := command.NewHostEcho(*message)
	utils.LogInfof("%s\n", cmd.String())
	*reply = command.EchoResponse{
		ID:      cmd.ID,
		Error:   "",
		Message: cmd.Body.(string),
	}
	return nil
}
