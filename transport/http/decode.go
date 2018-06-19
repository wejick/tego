//Copyright (c) 2015 Peter Bourgon
//Copyright (c) 2018 Gian Giovani

package http

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DecodeRequestFunc extracts a user-domain request object from an HTTP
// request object. It's designed to be used in HTTP servers, for server-side
// endpoints. One straightforward DecodeRequestFunc could be something that
// JSON decodes from the request body to the concrete request type.
type DecodeRequestFunc func(context.Context, *http.Request, httprouter.Params) (request interface{}, err error)

// NopRequestDecoder is a DecodeRequestFunc that can be used for requests that do not
// need to be decoded, and simply returns nil, nil.
func NopRequestDecoder(ctx context.Context, r *http.Request, _ httprouter.Params) (interface{}, error) {
	return nil, nil
}
