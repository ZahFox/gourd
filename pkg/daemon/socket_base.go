// +build !systemd

package daemon

import (
	"log"
	"net"
	"os"

	"github.com/zahfox/gourd/pkg/config"
)

// CreateListener creates a network listener to be used by gourdd
func CreateListener() (net.Listener, error) {
	socketPath := config.GetSocketPath()
	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatal(err)
	}

	return net.Listen("unix", socketPath)
}
