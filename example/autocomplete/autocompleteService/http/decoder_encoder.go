package http

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/tego/example/autocomplete/autocompleteService"
)

//ParameterDecoder get all request parameter from the request
func ParameterDecoder(ctx context.Context, r *http.Request, ps httprouter.Params) (request interface{}, err error) {
	query := r.URL.Query()

	request = autocompleteService.AutocompleteRequest{
		Keyword: query.Get("keyword"),
	}

	return
}
