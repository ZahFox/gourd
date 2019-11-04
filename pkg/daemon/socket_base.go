// +build !systemd

package daemon

import (
	"net"
	"os"

	"github.com/zahfox/gourd/pkg/config"
	"github.com/zahfox/gourd/pkg/utils"
)

// CreateListener creates a network listener to be used by gourdd
func CreateListener() (net.Listener, error) {
	socketPath := config.GetSocketPath()
	if err := os.RemoveAll(socketPath); err != nil {
		utils.LogFatal(err)
	}

	return net.Listen("unix", socketPath)
}
