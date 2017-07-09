package config

import "testing"

func TestHTTPUpstreamConfig_GetURL(t *testing.T) {
	type fields struct {
		Address string
		Port    string
		Schema  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "localhost",
			fields: fields{
				Schema:  "https://",
				Address: "localhost",
				Port:    "8080",
			},
			want: "https://localhost:8080",
		},
		{
			name: "empty",
			want: "",
		},
	}
	for _, tt := range tests {
		hcfg := &HTTPUpstreamConfig{
			Address: tt.fields.Address,
			Port:    tt.fields.Port,
			Schema:  tt.fields.Schema,
		}
		if got := hcfg.GetURL(); got != tt.want {
			t.Errorf("%q. HTTPUpstreamConfig.GetURL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSocketConfig_GetTarget(t *testing.T) {
	type fields struct {
		Address string
		Port    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "localhost",
			fields: fields{
				Address: "localhost",
				Port:    "8080",
			},
			want: "localhost:8080",
		},
		{
			name: "empty",
			want: "",
		},
	}
	for _, tt := range tests {
		scfg := &SocketConfig{
			Address: tt.fields.Address,
			Port:    tt.fields.Port,
		}
		if got := scfg.GetTarget(); got != tt.want {
			t.Errorf("%q. SocketConfig.GetTarget() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
