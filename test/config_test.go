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

	//generic
	if config.Get().Name != "smile" {
		t.Log("name failed")
		t.Fail()
	}
	if config.Get().Version != "0.0.1" {
		t.Log("version failed")
		t.Fail()
	}
	if config.Get().Description != "config example" {
		t.Log("description failed")
		t.Fail()
	}

	//db
	if config.Get().DB.Postgres["trackingDocument"].Host != "localhost:5432" {
		t.Log(config.Get().DB.Postgres["trackingDocument"].Host)
		t.Fail()
	}
	if config.Get().DB.Postgres["trackingDocument"].Database != "tracking" {
		t.Log("database failed")
		t.Fail()
	}
	if config.Get().DB.Postgres["trackingDocument"].SSLMode != "disable" {
		t.Log("ssl mode failed")
		t.Fail()
	}
	if config.Get().DB.Postgres["trackingDocument"].GetURL() != "postgres://gio:gio@localhost:5432/tracking?sslmode=disable" {
		t.Log("GetDSSN failed")
		t.Fail()
	}

	//http
	if config.Get().HTTP.Listen != "localhost" {
		t.Log("http listen failed")
		t.Fail()
	}
	if config.Get().HTTP.Port != "8080" {
		t.Log("http port failed")
		t.Fail()
	}

	//upstream
	if config.Get().Upstream.HTTP["orcinus"].Address != "localhost" {
		t.Log("http upstream Address failed")
		t.Fail()
	}
	if config.Get().Upstream.HTTP["orcinus"].Schema != "http://" {
		t.Log("http upstream Address failed")
		t.Fail()
	}
	if config.Get().Upstream.HTTP["orcinus"].Port != "" {
		t.Log("http upstream port failed")
		t.Fail()
	}
	if config.Get().Upstream.HTTP["orcinus"].GetURL() != "http://localhost" {
		t.Log("http upstream port failed")
		t.Fail()
	}
	if config.Get().Upstream.UnixSocket["docker"].FileDescriptor != "/var/log/docker.sock" {
		t.Log("upstream unix socket failed")
		t.Fail()
	}
	if config.Get().Upstream.Socket["gRPC"].Address != "localhost" {
		t.Log("socket upstream failed")
		t.Fail()
	}
	if config.Get().Upstream.Socket["gRPC"].Port != "8081" {
		t.Log("socket upstream failed")
		t.Fail()
	}
	if config.Get().Upstream.Socket["gRPC"].GetTarget() != "localhost:8081" {
		t.Log("socket upstream failed")
		t.Fail()
	}
}
