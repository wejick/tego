package config

type (
	//HTTPConfig http config
	HTTPConfig struct {
		Listen string `json:"listen"`
		Port   string `json:"port"`
	}
)
