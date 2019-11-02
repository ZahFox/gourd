package rpc

import (
	"log"

	"github.com/zahfox/gourd/pkg/command"
)

// Host accepts and responds to rpc commands for the local gourdd
type Host struct {
}

// Ping responds with pong
func (c *Host) Ping(_ interface{}, reply *string) error {
	cmd := command.NewHostPing()
	log.Printf("%s\n", cmd.String())
	*reply = "PONG"
	return nil
}

// Echo responds with the message that was sent to it
func (c *Host) Echo(message *string, reply *string) error {
	cmd := command.NewHostEcho(*message)
	log.Printf("%s\n", cmd.String())
	*reply = *message
	return nil
}
