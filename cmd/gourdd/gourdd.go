package main

import (
	"log"

	"github.com/zahfox/gourd/pkg/config"
	"github.com/zahfox/gourd/pkg/daemon"
)

func main() {
	config.Load()
	log.Printf("gourdd: listening for commands at: %s", daemon.GetSocketPath())
	daemon.GetDaemon().Listen()
}
