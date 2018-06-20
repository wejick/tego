package grpc

import (
	"context"

	"github.com/wejick/tego/endpoint"
)

// Server wraps an endpoint and implements http.Handler.
type Server struct {
	e               endpoint.Endpoint
	requestDecoder  DecodeRequestFunc
	responseEncoder EncodeResponseFunc
}

type serverOption func(*Server)

// New server instance which wrap and endpoint
func New(endpoint endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...serverOption) *Server {
	server := &Server{
		e:               endpoint,
		requestDecoder:  dec,
		responseEncoder: enc,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

// Handler which should be called from the gRPC binding of the service
// implementation. The incoming request parameter, and returned response
// parameter, are both gRPC types, not user-domain.
type Handler interface {
	ServeGRPC(context.Context, interface{}) (context.Context, interface{}, error)
}

// ServeGRPC implements Handler.ServeGRPC
func (s *Server) ServeGRPC(ctx context.Context, req interface{}) (retCtx context.Context, gRPCResponse interface{}, err error) {
	// TO DO supporting meta data

	request, err := s.requestDecoder(ctx, req)
	if err != nil {
		// TO DO : log and handle error here
		return
	}

	response, err := s.e(ctx, request)
	if err != nil {
		// TO DO : log and handle error here
		return
	}

	gRPCResponse, err = s.responseEncoder(ctx, response)
	if err != nil {
		// TO DO : log and handle error here
		return
	}

	return
}
