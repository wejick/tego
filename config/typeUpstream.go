package config

import "fmt"

type (
	//UpstreamConfig container
	UpstreamConfig struct {
		HTTP       map[string]*HTTPUpstreamConfig `json:"http"`
		UnixSocket map[string]UnixSocketConfig    `json:"unixSocket"`
		Socket     map[string]*SocketConfig       `json:"socket"`
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

	//SocketConfig tcp / udp upstream config
	SocketConfig struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	}
)

//GetURL gets from http upstream connection string
func (hcfg *HTTPUpstreamConfig) GetURL() string {
	if hcfg.Address == "" {
		return ""
	}
	if hcfg.Schema == "" {
		hcfg.Schema = "http"
	}
	if hcfg.Port == "" {
		return fmt.Sprintf("%s%s",
			hcfg.Schema,
			hcfg.Address,
		)
	}

	return fmt.Sprintf("%s%s:%s",
		hcfg.Schema,
		hcfg.Address,
		hcfg.Port,
	)
}

//GetTarget gets target address from upstream
func (scfg *SocketConfig) GetTarget() string {
	if scfg.Address == "" || scfg.Port == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s",
		scfg.Address,
		scfg.Port,
	)
}
