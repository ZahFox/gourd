package main

import (
	"github.com/zahfox/gourd/pkg/config"
	"github.com/zahfox/gourd/pkg/daemon"
	"github.com/zahfox/gourd/pkg/utils"
)

func main() {
	config.Load()
	utils.LogInfof("Running in a %s environment", config.EnvStr())
	utils.LogInfof("Listening for commands at %s", config.GetSocketPath())
	daemon.GetDaemon().Listen()
}
