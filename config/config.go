package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
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
	minifier := minify.New()
	minifier.AddFunc("text/javascript", js.Minify)
	f, err = minifier.Bytes("text/javascript", f)
	if err != nil {
		log.Println("Couldn't minify config file : " + configPath)
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
