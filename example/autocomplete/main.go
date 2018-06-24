package main

import (
	"log"
	"net"
	netHttp "net/http"

	goGRPC "google.golang.org/grpc"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/tego/example/autocomplete/autocompleteService"
	autocompleteGRPC "github.com/wejick/tego/example/autocomplete/autocompleteService/grpc"
	autocompleteHttp "github.com/wejick/tego/example/autocomplete/autocompleteService/http"
	"github.com/wejick/tego/transport/grpc"
	"github.com/wejick/tego/transport/http"
)

func main() {
	router := httprouter.New()

	gRPCServer := goGRPC.NewServer()

	autocompleteSvc := autocompleteService.New()
	autocompleteEndpoints := autocompleteService.NewEndpoints(autocompleteSvc)

	// http transport
	{
		suggestionServer := http.New(autocompleteEndpoints.Endpoints["suggestion"],
			autocompleteHttp.ParameterDecoder,
			http.EncodeJSONResponse)
		popularServer := http.New(autocompleteEndpoints.Endpoints["popular"],
			autocompleteHttp.ParameterDecoder,
			http.EncodeJSONResponse)

		router.GET("/suggestion", suggestionServer.HttprouterHandler)
		router.GET("/popular", popularServer.HttprouterHandler)
	}

	// grpc transport
	{
		suggestionServer := grpc.New(autocompleteEndpoints.Endpoints["suggestion"],
			autocompleteGRPC.RequestDecoder,
			autocompleteGRPC.SuggestionEncoder)
		popularServer := grpc.New(autocompleteEndpoints.Endpoints["popular"],
			autocompleteGRPC.RequestDecoder,
			autocompleteGRPC.PopularEncoder)
		autocompleteGRPC.RegisterAutocompleteServer(gRPCServer, autocompleteGRPC.MakeServerBinder(suggestionServer, popularServer))
	}

	sc, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("unable to listen: %+v", err)
	}
	defer gRPCServer.GracefulStop()

	go func() {
		err = gRPCServer.Serve(sc)
		if err != nil {
			log.Fatalf("unable to listen: %+v", err)
		}
	}()

	netHttp.ListenAndServe(":8080", router)
}
