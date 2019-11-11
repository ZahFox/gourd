// +build !systemd

package daemon

import (
	"net"
	"os"

	"github.com/zahfox/gourd/pkg/utils"
)

// CreateListener creates a network listener to be used by gourdd
func CreateListener(socketPath string) (net.Listener, error) {
	if err := os.RemoveAll(socketPath); err != nil {
		utils.LogFatal(err)
	}

	return net.Listen("unix", socketPath)
}
