package daemon

import (
	"log"
	"net"
	"os"

	"github.com/zahfox/gourd/pkg/config"
	"github.com/zahfox/gourd/pkg/daemon/rpc"
)

// Daemon is used to group together data related to gourdd
type Daemon struct {
	ID     uint8
	Socket net.Listener
}

// Listen will accept socket connections and forward them to the rpc command handler
func (d *Daemon) Listen() {
	defer d.Socket.Close()
	rpc.RegisterHandler()

	for {
		conn, err := d.Socket.Accept()
		if err != nil {
			log.Fatalf("Socket connection error: %+v\n", err)
		}

		log.Printf("New socket connection from %s\n", conn.RemoteAddr())
		rpc.HandleConnection(conn)
	}
}

var daemon *Daemon
var socketPath = ""

// GetDaemon will return the primary instance of gourdd
func GetDaemon() *Daemon {
	if daemon == nil {
		daemon = new(Daemon)
		daemon.ID = 0

		socket, err := CreateListener()
		if err != nil {
			log.Fatalf("Failed to listen to socket at %s. %s", GetSocketPath(), err)
		}

		daemon.Socket = socket
	}
	return daemon
}

// GetSocketPath returns the filesystem path to the command socket
func GetSocketPath() string {
	if socketPath != "" {
		return socketPath
	}

	path := os.Getenv("GOURD_GOURDD_SOCKET")
	if path != "" {
		socketPath = path
		return socketPath
	}

	env := config.GetEnv()
	if env == config.Debug {
		os.MkdirAll("/tmp/.gourd", 0700)
		socketPath = "/tmp/.gourd/gourdd-debug.sock"
	} else {
		socketPath = "/run/gourd/gourdd.sock"
	}

	return socketPath
}
