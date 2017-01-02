package config

type (
	//Config all config container
	Config struct {
		Name    string    `json:"name"`
		Version string    `json:"version"`
		DB      *DBConfig `json:"db"`
	}
)
