package config

import "testing"

func TestPostgresConfig_GetURL(t *testing.T) {
	tests := []struct {
		name string
		pCfg *PostgresConfig
		want string
	}{
		{
			name: "empty",
			pCfg: &PostgresConfig{},
			want: "",
		}, {
			name: "localhost",
			pCfg: &PostgresConfig{
				Database: "tracking",
				Host:     "localhost:5432",
				User:     "gio",
				Password: "gio",
				SSLMode:  "disable",
			},
			want: "postgres://gio:gio@localhost:5432/tracking?sslmode=disable",
		},
	}
	for _, tt := range tests {
		if got := tt.pCfg.GetURL(); got != tt.want {
			t.Errorf("%q. PostgresConfig.GetURL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
