package test

import (
	"context"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	nethttp "net/http"

	"github.com/stretchr/testify/assert"
	"github.com/wejick/tego/endpoint"
	"github.com/wejick/tego/transport/http"
)

func makeTestHaloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "halo", nil
	}
}

func TestServer(t *testing.T) {
	router := httprouter.New()
	server := http.New(makeTestHaloEndpoint(), http.NopRequestDecoder, http.EncodeJSONResponse)

	router.GET("/", server.HttprouterHandler)

	go nethttp.ListenAndServe("localhost:8080", router)

	resp, err := nethttp.Get("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	//Expected halo with eol
	expectedHalo, _ := json.Marshal("halo")
	eol, _ := hex.DecodeString("0a")
	expectedHalo = append(expectedHalo, eol...)

	assert.Equal(t, body, expectedHalo, "It's not halo", body)
}
