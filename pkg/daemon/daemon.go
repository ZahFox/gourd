package daemon

import (
	"log"
	"net"

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

		log.Printf("New socket connection from %s\n", conn.RemoteAddr().String())
		go rpc.HandleConnection(conn)
	}
}

var daemon *Daemon

// GetDaemon will return the primary instance of gourdd
func GetDaemon() *Daemon {
	if daemon == nil {
		daemon = new(Daemon)
		daemon.ID = 0

		socket, err := CreateListener()
		if err != nil {
			log.Fatalf("Failed to listen to socket at %s. %s", config.GetSocketPath(), err)
		}

		daemon.Socket = socket
	}
	return daemon
}
