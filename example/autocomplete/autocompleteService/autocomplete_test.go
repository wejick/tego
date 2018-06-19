package autocompleteService

import (
	"context"
	"reflect"
	"testing"
)

func Test_autocompleteService_Get(t *testing.T) {
	type args struct {
		ctx     context.Context
		request AutocompleteRequest
	}
	tests := []struct {
		name        string
		a           autocompleteService
		args        args
		wantRespond AutocompleteSuggestionRespond
		wantErr     bool
	}{
		{
			name: "empty",
			a:    autocompleteService{},
			args: args{
				ctx: nil,
				request: AutocompleteRequest{
					Keyword: "",
				},
			},
		},
		{
			name: "ba",
			a:    autocompleteService{},
			args: args{
				ctx: nil,
				request: AutocompleteRequest{
					Keyword: "ba",
				},
			},
			wantRespond: AutocompleteSuggestionRespond{
				Suggestions: []string{"baju", "baju bayi", "celana bayi"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := autocompleteService{}
			gotRespond, err := a.GetSuggestion(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("autocompleteService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespond, tt.wantRespond) {
				t.Errorf("autocompleteService.Get() = %v, want %v", gotRespond, tt.wantRespond)
			}
		})
	}
}

func Test_autocompleteService_GetPopular(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		a           autocompleteService
		args        args
		wantRespond AutocompletePopularRespond
		wantErr     bool
	}{
		{
			name: "default",
			a:    autocompleteService{},
			wantRespond: AutocompletePopularRespond{
				Popular: []string{"baju", "baju bayi"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := autocompleteService{}
			gotRespond, err := a.GetPopular(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("autocompleteService.GetPopular() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespond, tt.wantRespond) {
				t.Errorf("autocompleteService.GetPopular() = %v, want %v", gotRespond, tt.wantRespond)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want AutocompleteService
	}{
		{
			name: "default",
			want: autocompleteService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
