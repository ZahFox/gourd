// +build systemd

package daemon

import (
	"errors"
	"net"
	"os"
	"strconv"
)

// CreateListener creates a network listener to be used by gourdd
func CreateListener() (net.Listener, error) {
	if os.Getenv("LISTEN_PID") == strconv.Itoa(os.Getpid()) {
		f := os.NewFile(3, "socket")
		return net.FileListener(f)
	}

	return nil, errors.New("could not find any sockets supplied by systemd")
}
