package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
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

// HttprouterHandler implements httprouter.Handler so it can get Params data
func (s *Server) HttprouterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	request, err := s.requestDecoder(ctx, r, ps)
	if err != nil {
		// TO DO : log and handle error here
		return
	}

	response, err := s.e(ctx, request)
	if err != nil {
		// TO DO : log and handle error here
		return
	}

	err = s.responseEncoder(ctx, w, response)
	if err != nil {
		// TO DO : log and handle error here
		return
	}

}
