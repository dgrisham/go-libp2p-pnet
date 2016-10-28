package pnet

import (
	"io"

	mc "github.com/multiformats/go-multicodec"
	bmux "github.com/multiformats/go-multicodec/base/mux"
)

var (
	pathPSKv1   = []byte("/key/swarm/psk/1.0.0/")
	headerPSKv1 = mc.Header(pathPSKv1)
)

func decodeV1PSKKey(in io.Reader) (*[32]byte, error) {
	var err error
	in, err = mc.WrapTransformPathToHeader(in)
	if err != nil {
		return nil, err
	}
	err = mc.ConsumeHeader(in, headerPSKv1)
	if err != nil {
		return nil, err
	}

	in, err = mc.WrapTransformPathToHeader(in)
	if err != nil {
		return nil, err
	}

	out := [32]byte{}

	err = bmux.AllBasesMux().Decoder(in).Decode(out)
	return &out, err
}
