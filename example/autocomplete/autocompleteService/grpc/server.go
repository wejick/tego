package grpc

import (
	"context"

	"github.com/wejick/tego/example/autocomplete/autocompleteService/pb"
	"github.com/wejick/tego/transport/grpc"

	goGRPC "google.golang.org/grpc"
)

// ServerBinder Implement server interface from the protobuf definition
// it will be called by grpc server
// Why the name is server binder, it bind the tego server with the pb serve,
// hence the name is server binder
type ServerBinder struct {
	suggestionHandler grpc.Handler
	popularSuggestion grpc.Handler
}

// MakeServerBinder returns autocomplete service grpc server binder
func MakeServerBinder(suggestionHandler grpc.Handler, popularHandler grpc.Handler) ServerBinder {
	return ServerBinder{
		suggestionHandler: suggestionHandler,
		popularSuggestion: popularHandler,
	}
}

//RegisterAutocompleteServer register pb service to grpc server
func RegisterAutocompleteServer(grpcServer *goGRPC.Server, autocompleteServer ServerBinder) {
	pb.RegisterAutocompleteServer(grpcServer, autocompleteServer)
}

// GetSuggestion handler for suggestion service
func (s ServerBinder) GetSuggestion(ctx context.Context, req *pb.AutocompleteRequest) (*pb.SuggestionResponse, error) {
	_, resp, err := s.suggestionHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pb.SuggestionResponse)
	return response, err
}

// GetPopular handler for popular service
func (s ServerBinder) GetPopular(ctx context.Context, req *pb.AutocompleteRequest) (*pb.PopularResponse, error) {
	_, resp, err := s.popularSuggestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pb.PopularResponse)
	return response, err
}
