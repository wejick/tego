package config

import (
	"io/ioutil"
	"os"
	"testing"
)

type configTest struct {
	Version     int    `json:"version"`
	PackageName string `json:"name"`
}

func testLoadConfigFromFile(t *testing.T) {
	jsonTest := []byte("{'version':1,'name':'config'}")
	err := ioutil.WriteFile("config.json", jsonTest, os.ModePerm)
	if err != nil {
		t.Error("Couldn't write config.json")
	}

	var conf configTest
	err = LoadConfigFromFile("./config.json", conf)
	if err != nil {
		t.Error("Couldn't read config.json ", err)
	}

	if conf.Version != 1 {
		t.Error("Version expected 1, got ", conf.Version)
	}
	if conf.PackageName != "config" {
		t.Error("Package Name expected config, got ", conf.PackageName)
	}
}
