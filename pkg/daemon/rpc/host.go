package rpc

import (
	"fmt"
	"github.com/zahfox/gourd/pkg/command"
	"github.com/zahfox/gourd/pkg/distro"
	"github.com/zahfox/gourd/pkg/ltb"
	"github.com/zahfox/gourd/pkg/utils"
)

// Host accepts and responds to rpc commands for the local gourdd
type Host struct {
}

// Ping responds with pong
func (c *Host) Ping(_ interface{}, reply *command.PingResponse) error {
	cmd := command.NewHostPing()
	utils.LogInfo(cmd.String())
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
	utils.LogInfo(cmd.String())
	*reply = command.EchoResponse{
		ID:      cmd.ID,
		Error:   "",
		Message: cmd.Body.(string),
	}
	return nil
}

// Install responds with a message indicating whether the installation succeeded or failed
func (c *Host) Install(params *command.InstallRequestParams, reply *command.InstallResponse) error {
	i := params.Item
	cmd := command.NewHostInstall(i)
	utils.LogInfo(cmd.String())
	var msg string = "nothing happened"

	if i == "ltb" || i == "linux-toolbox" {
		if err := ltb.InstallForUser(params.User); err != nil {
			msg = fmt.Sprintf("linux-toolbox install failed\n%s", err.Error())
		} else {
			msg = fmt.Sprintf("successfully installed: %s", i)
		}
	} else {
		if err := distro.GetDistro().Install(i); err != nil {
			msg = fmt.Sprintf("distribution package install failed for: %s\n%s", i, err.Error())
		} else {
			msg = fmt.Sprintf("successfully installed: %s", i)
		}
	}

	*reply = command.InstallResponse{
		ID:      cmd.ID,
		Error:   "",
		Message: msg,
	}

	return nil
}
