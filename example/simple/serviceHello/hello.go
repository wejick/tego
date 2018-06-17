// Package serviceHello is service to echo any text thrown to it
// Steps on creating service
// 1. Define the service interface : HelloService
// 2. Create a private struct which implements the HelloService interface : helloService
// 3. Create initializer function returning the interface instance : New
// 4. Create your business handler function, in this case : Hello
// 4. Create endpoint function factory for Hello function : MakeHelloEndpoint
package serviceHello

import (
	"context"

	"github.com/wejick/tego/endpoint"
)

// HelloService is interface for this service
type HelloService interface {
	Hello(c context.Context, text string) (string, error)
}

type helloService struct{}

// New Creating new HelloService instance
func New() HelloService {
	return helloService{}
}

// Hello is function to handle business logic behind hello endpoint
func (h helloService) Hello(c context.Context, text string) (output string, err error) {
	if text != "" {
		output = text
	} else {
		output = "Hello ;)"
	}

	return
}

// MakeHelloEndpoint create helloService's hello endpoint
func MakeHelloEndpoint(helloService HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if request == nil {
			request = ""
		}
		return helloService.Hello(ctx, request.(string))
	}
}
