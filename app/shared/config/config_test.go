package config

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ivzb/achievers_server/app/shared/database"
	"github.com/ivzb/achievers_server/app/shared/token"
)

func TestNewConfig_ValidBytes(t *testing.T) {
	expectedConfig := &Config{
		Database: database.Info{
			Type: "Postgree",
			Postgre: database.PostgreInfo{
				Username:  "Admin",
				Password:  "",
				Name:      "achievers",
				Hostname:  "127.0.0.1",
				Port:      5432,
				Parameter: " sslmode=disable",
			},
		},
		Token: token.Info{
			Path: "rsa/token.pem",
		},
	}

	c, err := json.Marshal(expectedConfig)

	if err != nil {
		t.Fatalf("Config marshal error: %v", err)
	}

	actualConfig, err := New(c)

	if err != nil {
		t.Fatalf("New config returned error: %v", err)
	}

	if !cmp.Equal(expectedConfig, actualConfig) {
		t.Fatalf("Config returned unexpected value:\nexpected %#v,\nactual %#v",
			expectedConfig, actualConfig)
	}
}

func TestNewConfig_EmptyBytes(t *testing.T) {
	var bytes []byte

	_, err := New(bytes)

	if err == nil {
		t.Fatalf("Config should returned error")
	}
}

func TestNewConfig_InvalidBytes(t *testing.T) {
	bytes := []byte("random string")

	_, err := New(bytes)

	if err == nil {
		t.Fatalf("Config should returned error")
	}
}
