package main

import (
	netHttp "net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/tego/example/autocomplete/autocompleteService"
	autocompleteHttp "github.com/wejick/tego/example/autocomplete/autocompleteService/http"
	"github.com/wejick/tego/transport/http"
)

func main() {
	autocompleteSvc := autocompleteService.New()
	autocompleteEndpoints := autocompleteService.NewEndpoints(autocompleteSvc)

	suggestionServer := http.New(autocompleteEndpoints.Endpoints["suggestion"],
		autocompleteHttp.ParameterDecoder,
		http.EncodeJSONResponse)
	popularServer := http.New(autocompleteEndpoints.Endpoints["popular"],
		autocompleteHttp.ParameterDecoder,
		http.EncodeJSONResponse)

	router := httprouter.New()

	router.GET("/suggestion", suggestionServer.HttprouterHandler)
	router.GET("/popular", popularServer.HttprouterHandler)

	netHttp.ListenAndServe(":8080", router)
}
