package config

type (
	//DBConfig database config container
	DBConfig struct {
		Postgres map[string]PostgresConfig `json:"postgres"`
	}

	//PostgresConfig postgress config
	PostgresConfig struct {
		Database string `json:"database"`
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)
