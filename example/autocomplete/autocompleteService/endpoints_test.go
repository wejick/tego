package autocompleteService

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wejick/tego/endpoint"
)

func TestMakeSuggestionEndpoint(t *testing.T) {
	service := autocompleteService{}
	endpoint := MakeSuggestionEndpoint(service)

	endpointResult, err := endpoint(nil, AutocompleteRequest{Keyword: "baju"})
	if err != nil {
		t.Error("error occurend", err)
	}
	directResult, _ := service.GetSuggestion(nil, AutocompleteRequest{Keyword: "baju"})
	assert.Equal(t, endpointResult.(AutocompleteSuggestionRespond), directResult)

	//test nil
	endpointResult, err = endpoint(nil, nil)
	if err != nil {
		t.Error("error occurend", err)
	}
	directResult, _ = service.GetSuggestion(nil, AutocompleteRequest{})
	assert.Equal(t, endpointResult.(AutocompleteSuggestionRespond), directResult)
}

func TestMakePopularEndpoint(t *testing.T) {
	service := autocompleteService{}
	endpoint := MakePopularEndpoint(service)

	endpointResult, err := endpoint(nil, AutocompleteRequest{})
	if err != nil {
		t.Error("error occurend", err)
	}
	directResult, _ := service.GetPopular(nil)
	assert.Equal(t, endpointResult.(AutocompletePopularRespond), directResult)
}

func TestNewEndpoints(t *testing.T) {
	service := autocompleteService{}
	set := NewEndpoints(service)

	naiveSet := endpoint.Set{}
	naiveSet.Endpoints = make(map[string]endpoint.Endpoint)
	naiveSet.Endpoints["suggestion"] = MakeSuggestionEndpoint(service)
	naiveSet.Endpoints["popular"] = MakePopularEndpoint(service)

	naiveSuggestionResult, _ := naiveSet.Endpoints["suggestion"](nil, AutocompleteRequest{Keyword: "baju"})
	suggestionResult, _ := set.Endpoints["suggestion"](nil, AutocompleteRequest{Keyword: "baju"})
	assert.Equal(t, naiveSuggestionResult, suggestionResult)

	naivePopularResult, _ := naiveSet.Endpoints["popular"](nil, AutocompleteRequest{})
	popularResult, _ := set.Endpoints["popular"](nil, AutocompleteRequest{})
	assert.Equal(t, naivePopularResult, popularResult)
}
