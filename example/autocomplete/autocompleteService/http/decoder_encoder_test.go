package http

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/tego/example/autocomplete/autocompleteService"
)

func TestParameterDecoder(t *testing.T) {
	emptyReq, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	bajuReq, _ := http.NewRequest("GET", "http://localhost:8080/?keyword=baju", nil)
	type args struct {
		ctx     context.Context
		request *http.Request
		ps      httprouter.Params
	}
	tests := []struct {
		name        string
		args        args
		wantRequest interface{}
		wantErr     bool
	}{
		{
			name: "empty",
			args: args{
				request: emptyReq,
			},
			wantRequest: autocompleteService.AutocompleteRequest{},
		},
		{
			name: "empty",
			args: args{
				request: bajuReq,
			},
			wantRequest: autocompleteService.AutocompleteRequest{
				Keyword: "baju",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRequest, err := ParameterDecoder(tt.args.ctx, tt.args.request, tt.args.ps)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParameterDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRequest, tt.wantRequest) {
				t.Errorf("ParameterDecoder() = %v, want %v", gotRequest, tt.wantRequest)
			}
		})
	}
}
