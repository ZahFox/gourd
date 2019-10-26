package daemon

import (
	"io"
	"log"
	"net"
)

// Daemon is used to group together data related to gourdd
type Daemon struct {
	ID     string
	Socket net.Listener
}

// Listen will accept socket connections and respond to their commands
func (d *Daemon) Listen() {
	defer d.Socket.Close()

	for {
		conn, err := d.Socket.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
}

var daemon Daemon

func init() {
	daemon.ID = "1"
	socket, err := CreateListener()

	if err != nil {
		log.Fatalf("failed to listen to socket at %s. %s", GetSocketPath(), err)
	}

	daemon.Socket = socket
}

// GetDaemon will return the primary instance of gourdd
func GetDaemon() *Daemon {
	return &daemon
}

// GetSocketPath returns the filesystem path to the command socket
func GetSocketPath() string {
	return "/run/gourd/gourdd.sock"
}

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	io.Copy(c, c)
	c.Close()
}
