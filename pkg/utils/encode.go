package utils

import (
	"github.com/ugorji/go/codec"
)

var cborHandle *codec.CborHandle

func init() {
	cborHandle = new(codec.CborHandle)
	cborHandle.ErrorIfNoField = false
	cborHandle.WriterBufferSize = 8192
	cborHandle.ReaderBufferSize = 8192
}

// CborEncoder will encode data in cbor format
func CborEncoder(buffer *[]byte) *codec.Encoder {
	return codec.NewEncoderBytes(buffer, cborHandle)
}

// CborDecoder will decode data from a cbor format
func CborDecoder(buffer *[]byte) *codec.Decoder {
	return codec.NewDecoderBytes(*buffer, cborHandle)
}
