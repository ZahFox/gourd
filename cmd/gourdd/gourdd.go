package main

import (
	"log"

	"github.com/zahfox/gourd/pkg/daemon"
)

func main() {
	log.Printf("gourdd: listening for commands at: %s", daemon.GetSocketPath())
	daemon.GetDaemon().Listen()
}
