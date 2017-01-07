package config

import "fmt"

type (
	//UpstreamConfig container
	UpstreamConfig struct {
		HTTP       map[string]*HTTPUpstreamConfig `json:"http"`
		UnixSocket map[string]UnixSocketConfig    `json:"unixSocket"`
	}

	//HTTPUpstreamConfig http upstream config
	HTTPUpstreamConfig struct {
		Address string `json:"address"`
		Port    string `json:"port"`
		Schema  string `json:"schema"`
	}

	//UnixSocketConfig unix socket config
	UnixSocketConfig struct {
		FileDescriptor string `json:"fileDescriptor"`
	}
)

//GetURL gets dssn connection string
func (hcfg *HTTPUpstreamConfig) GetURL() string {
	return fmt.Sprintf("%s%s:%s",
		hcfg.Schema,
		hcfg.Address,
		hcfg.Port,
	)
}
