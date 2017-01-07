package config

import "fmt"

type (
	//UpstreamConfig container
	UpstreamConfig struct {
		HTTP map[string]*HTTPUpstreamConfig `json:"http"`
	}

	//HTTPUpstreamConfig config
	HTTPUpstreamConfig struct {
		Address string `json:"address"`
		Port    string `json:"Port"`
		Schema  string `json:"Schema"`
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
