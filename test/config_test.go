package configTest

import (
	"testing"

	"github.com/wejick/tego/config"
)

func TestLoadConfigFromFile(t *testing.T) {
	err := config.LoadConfigFromFile("./config.json")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if config.Get().Name != "smile" {
		t.Log("name failed")
		t.Fail()
	}
	if config.Get().DB.Postgres["trackingDocument"].Host != "localhost:5432" {
		t.Log(config.Get().DB.Postgres["trackingDocument"].Host)
		t.Fail()
	}
}
