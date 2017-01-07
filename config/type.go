package config

type (
	//Config all config container
	Config struct {
		Name        string         `json:"name"`
		Version     string         `json:"version"`
		Description string         `json:"description"`
		DB          *DBConfig      `json:"db"`
		HTTP        *HTTPConfig    `json:"http"`
		Upstream    UpstreamConfig `json:"upstream"`
	}
)
