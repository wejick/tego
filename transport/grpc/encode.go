//Copyright (c) 2015 Peter Bourgon
//Copyright (c) 2018 Gian Giovani

package grpc

import (
	"context"
)

// EncodeResponseFunc encodes the passed response object to the gRPC response
// writer. It's designed to be used in gRPC servers, for server-side
// endpoints.
type EncodeResponseFunc func(context.Context, interface{}) (interface{}, error)
