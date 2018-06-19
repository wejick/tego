package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wejick/tego/example/simple/serviceHello"
	"github.com/wejick/tego/transport/http"
	grace "gopkg.in/paytm/grace.v1"
)

func main() {
	hellosvc := serviceHello.New()
	helloServer := http.New(serviceHello.MakeHelloEndpoint(hellosvc), http.NopRequestDecoder, http.EncodeJSONResponse)

	router := httprouter.New()
	router.GET("/", helloServer.HttprouterHandler)

	grace.Serve(":8080", router)
}
