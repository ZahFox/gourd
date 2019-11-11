// +build !systemd

package daemon

import (
	"net"
	"os"

	"github.com/zahfox/gourd/pkg/utils"
)

// CreateListener creates a network listener to be used by gourdd
func CreateListener(socketPath string) (net.Listener, error) {
	gourdID, err := utils.GetGourdID()
	if err != nil {
		return nil, nil
	}

	if err = os.RemoveAll(socketPath); err != nil {
		utils.LogFatal(err)
	}

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return l, err
	}

	os.Chown(socketPath, gourdID.UID, gourdID.GID)
	os.Chmod(socketPath, 0660)
	return l, nil
}
