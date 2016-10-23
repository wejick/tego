package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

/*
LoadConfigFromFile load json configuration file from configPath
to given struct
*/
func LoadConfigFromFile(configPath string, cfg interface{}) (err error) {
	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		errorLog := "Couldn't read config file : " + configPath + "" + err.Error()
		log.Println(errorLog)
		return errors.New(errorLog)
	}
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		errorLog := "Couldn't read config file : " + configPath + "" + err.Error()
		log.Println(errorLog)
		return errors.New(errorLog)
	}

	return
}
