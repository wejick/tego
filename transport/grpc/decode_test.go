//Copyright (c) 2015 Peter Bourgon
//Copyright (c) 2018 Gian Giovani

package grpc

import (
	"context"
	"reflect"
	"testing"
)

func TestNopRequestDecoder(t *testing.T) {
	type args struct {
		ctx context.Context
		i   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: nil,
		},
		{
			name: "not empty",
			args: args{
				i: "not empty",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NopRequestDecoder(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("NopRequestDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NopRequestDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
