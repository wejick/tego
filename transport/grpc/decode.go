//Copyright (c) 2015 Peter Bourgon
//Copyright (c) 2018 Gian Giovani

package grpc

import (
	"context"
)

// DecodeRequestFunc extracts a user-domain request object from an gRPC
// request object. It's designed to be used in gRPC servers, for server-side
// endpoints.
type DecodeRequestFunc func(context.Context, interface{}) (interface{}, error)

// NopRequestDecoder is a DecodeRequestFunc that can be used for requests that do not
// need to be decoded, and simply returns nil, nil.
func NopRequestDecoder(ctx context.Context, i interface{}) (interface{}, error) {
	return nil, nil
}
