package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var cfg *Config

/*
LoadConfigFromFile load json configuration file from configPath
to given struct
*/
func LoadConfigFromFile(configPath string) (err error) {
	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println("Couldn't read config file : " + configPath)
		return
	}
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		log.Println("Couldn't read config file : " + configPath)
		return
	}

	return
}

//Get returns config container pointer
func Get() *Config {
	return cfg
}
