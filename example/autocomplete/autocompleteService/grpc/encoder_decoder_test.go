package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/wejick/tego/example/autocomplete/autocompleteService"
	"github.com/wejick/tego/example/autocomplete/autocompleteService/pb"
)

func TestRequestDecoder(t *testing.T) {
	type args struct {
		ctx context.Context
		req interface{}
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
				req: pb.AutocompleteRequest{},
			},
			wantRequest: autocompleteService.AutocompleteRequest{},
		},
		{
			name: "baju",
			args: args{
				req: pb.AutocompleteRequest{
					Keyword: "baju",
				},
			},
			wantRequest: autocompleteService.AutocompleteRequest{
				Keyword: "baju",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRequest, err := RequestDecoder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRequest, tt.wantRequest) {
				t.Errorf("RequestDecoder() = %v, want %v", gotRequest, tt.wantRequest)
			}
		})
	}
}

func TestSuggestionEncoder(t *testing.T) {
	type args struct {
		ctx      context.Context
		response interface{}
	}
	tests := []struct {
		name             string
		args             args
		wantGrpcResponse interface{}
		wantErr          bool
	}{
		{
			name: "empty",
			args: args{
				response: autocompleteService.AutocompleteSuggestionRespond{},
			},
			wantGrpcResponse: pb.SuggestionResponse{},
		},
		{
			name: "baju",
			args: args{
				response: autocompleteService.AutocompleteSuggestionRespond{
					Suggestions: []string{"baju"},
				},
			},
			wantGrpcResponse: pb.SuggestionResponse{
				Suggestions: []string{"baju"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrpcResponse, err := SuggestionEncoder(tt.args.ctx, tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("SuggestionEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGrpcResponse, tt.wantGrpcResponse) {
				t.Errorf("SuggestionEncoder() = %v, want %v", gotGrpcResponse, tt.wantGrpcResponse)
			}
		})
	}
}

func TestPopularEncoder(t *testing.T) {
	type args struct {
		ctx      context.Context
		response interface{}
	}
	tests := []struct {
		name             string
		args             args
		wantGrpcResponse interface{}
		wantErr          bool
	}{
		{
			name: "empty",
			args: args{
				response: autocompleteService.AutocompletePopularRespond{},
			},
			wantGrpcResponse: pb.PopularResponse{},
		},
		{
			name: "baju",
			args: args{
				response: autocompleteService.AutocompletePopularRespond{
					Popular: []string{"baju"},
				},
			},
			wantGrpcResponse: pb.PopularResponse{
				Popular: []string{"baju"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrpcResponse, err := PopularEncoder(tt.args.ctx, tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopularEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGrpcResponse, tt.wantGrpcResponse) {
				t.Errorf("PopularEncoder() = %v, want %v", gotGrpcResponse, tt.wantGrpcResponse)
			}
		})
	}
}
