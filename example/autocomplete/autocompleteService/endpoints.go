package autocompleteService

import (
	"context"

	"github.com/wejick/tego/endpoint"
)

//MakeSuggestionEndpoint create suggestion endpoint
func MakeSuggestionEndpoint(service AutocompleteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (responds interface{}, err error) {
		var req AutocompleteRequest
		if request != nil {
			req = request.(AutocompleteRequest)
		}
		responds, err = service.GetSuggestion(ctx, req)
		return
	}
}

//MakePopularEndpoint create autocomplete endpoint
func MakePopularEndpoint(service AutocompleteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (responds interface{}, err error) {
		responds, err = service.GetPopular(ctx)
		return
	}
}

//NewEndpoints returns Set of AutocompleteService endpoint with all default middleware
func NewEndpoints(service AutocompleteService) (set endpoint.Set) {
	set.Endpoints = make(map[string]endpoint.Endpoint)

	//make suggestion endpoint
	var suggestionEndpoint endpoint.Endpoint
	{
		suggestionEndpoint = MakeSuggestionEndpoint(service)
	}

	var popularEndpoint endpoint.Endpoint
	{
		popularEndpoint = MakePopularEndpoint(service)
	}

	set.Endpoints["suggestion"] = suggestionEndpoint
	set.Endpoints["popular"] = popularEndpoint

	return
}
