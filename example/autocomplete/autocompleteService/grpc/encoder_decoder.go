package grpc

import (
	"context"

	"github.com/wejick/tego/example/autocomplete/autocompleteService"
	"github.com/wejick/tego/example/autocomplete/autocompleteService/pb"
)

//RequestDecoder request into autocompleteRequest
func RequestDecoder(ctx context.Context, req interface{}) (request interface{}, err error) {
	var r pb.AutocompleteRequest
	if req != nil {
		r = req.(pb.AutocompleteRequest)
	}
	request = autocompleteService.AutocompleteRequest{
		Keyword: r.Keyword,
	}
	return
}

//SuggestionEncoder encode suggestionResponse to pb message
func SuggestionEncoder(ctx context.Context, response interface{}) (grpcResponse interface{}, err error) {
	resp := response.(autocompleteService.AutocompleteSuggestionRespond)
	grpcResponse = pb.SuggestionResponse{
		Suggestions: resp.Suggestions,
	}
	return
}

//PopularEncoder encode popularResponse to pb message
func PopularEncoder(ctx context.Context, response interface{}) (grpcResponse interface{}, err error) {
	resp := response.(autocompleteService.AutocompletePopularRespond)
	grpcResponse = pb.PopularResponse{
		Popular: resp.Popular,
	}
	return
}
