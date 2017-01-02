package config

import "fmt"

type (
	//DBConfig database config container
	DBConfig struct {
		Postgres map[string]*PostgresConfig `json:"postgres"`
	}

	//PostgresConfig postgress config
	PostgresConfig struct {
		Database string `json:"database"`
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		SSLMode  string `json:"sslmode"`
	}
)

//GetURL gets dssn connection string
func (pCfg *PostgresConfig) GetURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		pCfg.User,
		pCfg.Password,
		pCfg.Host,
		pCfg.Database,
		pCfg.SSLMode,
	)
}
