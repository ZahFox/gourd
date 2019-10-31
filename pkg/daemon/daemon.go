package daemon

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/zahfox/gourd/pkg/config"
	cmd "github.com/zahfox/gourd/pkg/daemon/rpc"
)

// Daemon is used to group together data related to gourdd
type Daemon struct {
	ID     string
	Socket net.Listener
}

// Listen will accept socket connections and respond to their commands
func (d *Daemon) Listen() {
	defer d.Socket.Close()
	http.Serve(d.Socket, nil)
}

var daemon *Daemon
var socketPath = ""

// GetDaemon will return the primary instance of gourdd
func GetDaemon() *Daemon {
	if daemon == nil {
		daemon = new(Daemon)
		daemon.ID = "1"

		commandHandler := new(cmd.CommandHandler)
		rpc.Register(commandHandler)
		rpc.HandleHTTP()

		socket, err := CreateListener()
		if err != nil {
			log.Fatalf("failed to listen to socket at %s. %s", GetSocketPath(), err)
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

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	io.Copy(c, c)
	c.Close()
}
