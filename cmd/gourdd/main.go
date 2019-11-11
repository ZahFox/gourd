package main

import (
	"github.com/zahfox/gourd/internal/gourdd/config"
	"github.com/zahfox/gourd/pkg/daemon"
	"github.com/zahfox/gourd/pkg/utils"
)

func main() {
	socketPath := config.GetSocketPath()
	utils.LogInfof("Running in a %s environment", utils.EnvStr())
	utils.LogInfof("Listening for commands at %s", socketPath)
	daemon.GetDaemon(socketPath).Listen()
}
