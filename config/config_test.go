package config

import (
	"reflect"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name: "get from ../test",
			args: args{
				configPath: "../test/config.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := LoadConfigFromFile(tt.args.configPath); (err != nil) != tt.wantErr {
			t.Errorf("%q. LoadConfigFromFile() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "empty",
			want: Get(),
		},
	}
	for _, tt := range tests {
		if got := Get(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
