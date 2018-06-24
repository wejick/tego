package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wejick/tego/example/autocomplete/autocompleteService/pb"
	goGRPC "google.golang.org/grpc"
)

type okSuggestion struct{}
type okPopular struct{}

func (o okSuggestion) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	if req.(*pb.AutocompleteRequest) == nil {
		return nil, nil, errors.New("error")
	}
	return nil, &pb.SuggestionResponse{Suggestions: []string{"OK"}}, nil
}

func (o okPopular) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	if req.(*pb.AutocompleteRequest) == nil {
		return nil, nil, errors.New("error")
	}
	return nil, &pb.PopularResponse{Popular: []string{"OK"}}, nil
}

func TestMakeServerBinder(t *testing.T) {
	// empty
	empty := ServerBinder{
		suggestionHandler: nil,
		popularSuggestion: nil,
	}
	assert.Equal(t, empty, MakeServerBinder(nil, nil))

	// ok no error
	okBinder := MakeServerBinder(okSuggestion{}, okPopular{})
	resp, _ := okBinder.GetSuggestion(nil, &pb.AutocompleteRequest{})
	assert.Equal(t, []string{"OK"}, resp.Suggestions)

	respPop, _ := okBinder.GetPopular(nil, &pb.AutocompleteRequest{})
	assert.Equal(t, []string{"OK"}, respPop.Popular)

	// ok error
	resp, err := okBinder.GetSuggestion(nil, nil)

	assert.Nil(t, resp)
	assert.Equal(t, errors.New("error"), err)

	respPop, err = okBinder.GetPopular(nil, nil)
	assert.Nil(t, respPop)
	assert.Equal(t, errors.New("error"), err)
}

func TestRegisterAutocompleteServer(t *testing.T) {
	server := goGRPC.NewServer()
	okBinder := MakeServerBinder(okSuggestion{}, okPopular{})
	RegisterAutocompleteServer(server, okBinder)

	assert.Equal(t, len(server.GetServiceInfo()), 1)
	assert.Equal(t, len(server.GetServiceInfo()["pb.autocomplete"].Methods), 2)
}
