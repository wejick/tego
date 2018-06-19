# Autocomplete API example
We will build a service which give us two functionality :
1. Get suggestion word based on given keyword.
1. Get current popular keyword.

Of course the data will be dummy, but we will get the idea.

In general the building block of Tego is 3 things :
1. Service      : Where the business code lays.
1. Endpoint     : Where the business logic get exposed to the outside world interface (eg grpc, http).
1. transport    : The way the outside world access the endpoint.


On building autocomplete API we will have :
- 1 service : The Autocomplete Service
- which has 2 endpoints : 
  1. get suggestion : accept `keyword` parameter and return list of matched keyword
  1. get popular : doesn't need any parameter and return list of popular keyword

- will support 1 transport : http which will return responds in JSON.

## Building the service

Building service started by defining the base interface, in our case we will have two functionality :

```golang
type AutocompleteService interface {
	GetSuggestion(context.Context, AutocompleteRequest) (AutocompleteSuggestionRespond, error)
	GetPopular(context.Context) (AutocompletePopularRespond, error)
}
```

We can see that we need 3 other interface for the request and responds of each function.

```golang
type AutocompleteRequest struct {
	Keyword string
}

type AutocompleteSuggestionRespond struct {
	Suggestions []string `json:"suggestion"`
	TookTime    int      `json:"took_time"`
}

type AutocompletePopularRespond struct {
	Popular  []string `json:"popular"`
	TookTime int      `json:"took_time"`
}
```

Starting from here we can implement the logic.

## Creating the endpoint
Endpoint defined by this type :
```golang
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
```
It will be the the entrypoint for any RPC method we choose to interact with the services logic. That's why we can see it generalize the parameter and response as interface.

We create the endpoint with a "factory" faction :
```golang
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
```

## About request decoding and response encoding
For some avid developers will notice that our service function has defined struct as parameter and response, it means any incoming request from outside the world have to be mapped to that struct. That's where the decoding and encoding function come, they are difference for each transport method.

Get popular doesn't expect any parameter, for get suggestion http interface we will have the parameter comes as query string like : `?keyword=baju`. So we have this request decoder :
```golang
func ParameterDecoder(ctx context.Context, r *http.Request, ps httprouter.Params) (request interface{}, err error) {
	query := r.URL.Query()

	request = autocompleteService.AutocompleteRequest{
		Keyword: query.Get("keyword"),
	}

	return
}
```

And since both are responding in JSON so we can use default JSON encoder from tego.

## Wiring all the stuffs

Now we need to wire the service, endpoint and transport into working app, that's where the `server` comes.

If you have multiple endpoint on your service, you can use Set to wrap all of your endpoint for ergonomy. In the NewEndpoints function we wrap our endpoints in the Set.

```golang
autocompleteSvc := autocompleteService.New()
	autocompleteEndpoints := autocompleteService.NewEndpoints(autocompleteSvc)
```

Here is where we create server to wire all the components, notice that we use previously created decoder and tego provided JSON encoder.
```golang
	suggestionServer := http.New(autocompleteEndpoints.Endpoints["suggestion"],
		autocompleteHttp.ParameterDecoder,
        http.EncodeJSONResponse)
```

Last step is to route http request to our server. Tego are opinionated using httprouter as default muxer.
```golang
router := httprouter.New()

router.GET("/suggestion", suggestionServer.HttprouterHandler)
router.GET("/popular", popularServer.HttprouterHandler)

netHttp.ListenAndServe(":8080", router)
```

Now compile and run our brand new autocomplete service