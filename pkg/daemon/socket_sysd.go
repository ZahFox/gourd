// +build systemd

package daemon

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "gourdd: Hello, Client!\n")
}

func pong(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "PONG\n")
}

// MakeSocket creates the Unix domain socket used for sending commands to gourdd
func MakeSocket() {
	if os.Getenv("LISTEN_PID") == strconv.Itoa(os.Getpid()) {
		f := os.NewFile(3, "socket")
		l, err := net.FileListener(f)

		if err != nil {
			log.Fatal(err)
		}

		http.HandleFunc("/", handleRequest)
		http.HandleFunc("/ping", pong)
		http.Serve(l, nil)
	} else {
		log.Panicf("gourdd: could not find any sockets supplied by systemd")
	}
}
