package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_ServeGRPC(t *testing.T) {
	assert.Equal(t, nil, nil)

	//test empty
	serviceEmpty := New(nil, nil, nil)
	assert.Equal(t, &Server{
		e:               nil,
		requestDecoder:  nil,
		responseEncoder: nil,
	}, serviceEmpty)

	//test dummy service
	emptyEndpoint := func(ctx context.Context, res interface{}) (response interface{}, err error) {
		return res, nil
	}
	decoder := func(ctx context.Context, i interface{}) (interface{}, error) {
		return i, nil
	}
	encoder := func(ctx context.Context, i interface{}) (interface{}, error) { return i, nil }
	serviceDummy := New(emptyEndpoint, decoder, encoder)
	serviceDummyNaive := &Server{
		e:               emptyEndpoint,
		requestDecoder:  decoder,
		responseEncoder: encoder,
	}
	decoderVal, _ := serviceDummy.requestDecoder(nil, "baju")
	decoderNaiveVal, _ := serviceDummyNaive.requestDecoder(nil, "baju")
	assert.Equal(t, decoderNaiveVal, decoderVal)

	decoderVal, _ = serviceDummy.responseEncoder(nil, "baju")
	decoderNaiveVal, _ = serviceDummyNaive.requestDecoder(nil, "baju")
	assert.Equal(t, decoderNaiveVal, decoderVal)

	_, res, _ := serviceDummy.ServeGRPC(context.Background(), "baju")
	assert.Equal(t, "baju", res)
}
