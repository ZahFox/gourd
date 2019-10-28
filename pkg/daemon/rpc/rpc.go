package rpc

import "errors"

type CommandHandler struct {
}

func (c CommandHandler) Handle(command *string, reply *string) error {
	switch *command {
	case "PING":
		*reply = "PONG"
		return nil
	default:
		return errors.New("invalid command type")
	}
}
