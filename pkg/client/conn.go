package client

import (
	"net"
	"net/rpc"

	"github.com/ugorji/go/codec"
	"github.com/zahfox/gourd/pkg/config"
	"github.com/zahfox/gourd/pkg/utils"
)

var handle *codec.CborHandle

func init() {
	handle = new(codec.CborHandle)
	handle.WriterBufferSize = 8192
	handle.ReaderBufferSize = 8192
}

func getConn() *rpc.Client {
	nc, err := net.Dial("unix", config.GetSocketPath())
	if err != nil {
		utils.LogFatal("Failed to connect to gourdd. ", err)
	}

	rpcCodec := codec.GoRpc.ClientCodec(nc, handle)
	return rpc.NewClientWithCodec(rpcCodec)
}
