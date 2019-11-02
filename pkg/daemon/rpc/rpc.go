package rpc

import (
	"io"
	"net/rpc"

	"github.com/ugorji/go/codec"
)

var registered = false

var handler *codec.CborHandle

// HandleConnection creates a new gourdd rpc codec
func HandleConnection(conn io.ReadWriteCloser) {
	rpc.ServeCodec(codec.GoRpc.ServerCodec(conn, handler))
}

// RegisterHandler initializes the rpc command handler
func RegisterHandler() {
	if registered {
		return
	}

	registered = true
	handler = new(codec.CborHandle)
	handler.WriterBufferSize = 8192
	handler.ReaderBufferSize = 8192
	host := new(Host)
	rpc.Register(host)
}
