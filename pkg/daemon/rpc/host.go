package rpc

import (
	"log"
	"os"

	"github.com/zahfox/gourd/pkg/command"
)

// Host accepts and responds to rpc commands for the local gourdd
type Host struct {
}

// TODO: Remove this and use a logging service instead
var sl = log.New(os.Stdout, "", 0)

// Ping responds with pong
func (c *Host) Ping(_ interface{}, reply *string) error {
	cmd := command.NewHostPing()
	sl.Printf("%s\n", cmd.String())
	*reply = "PONG"
	return nil
}

// Echo responds with the message that was sent to it
func (c *Host) Echo(message *string, reply *string) error {
	cmd := command.NewHostEcho(*message)
	sl.Printf("%s\n", cmd.String())
	*reply = *message
	return nil
}
