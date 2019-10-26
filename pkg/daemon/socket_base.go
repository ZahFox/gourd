// +build !systemd

package daemon

import (
	"log"
	"net"
	"os"
)

// CreateListener creates a network listener to be used by gourdd
func CreateListener() (net.Listener, error) {
	socketPath := GetSocketPath()
	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatal(err)
	}

	return net.Listen("unix", socketPath)
}
