package config

type (
	//UpstreamConfig container
	UpstreamConfig struct {
		HTTP map[string]HTTPUpstreamConfig `json:"http"`
	}

	//HTTPUpstreamConfig config
	HTTPUpstreamConfig struct {
		Address string `json:"address"`
		Port    string `json:"Port"`
		Schema  string `json:"Schema"`
	}
)
