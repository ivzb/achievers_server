package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ivzb/achievers_server/app/shared/database"
	"github.com/ivzb/achievers_server/app/shared/token"
)

func TestNewConfig_ValidBytes(t *testing.T) {
	expectedConfig := &Config{
		Database: database.Info{
			Type: "MySQL",
			MySQL: database.MySQLInfo{
				Username:  "root",
				Password:  "",
				Name:      "achievers",
				Hostname:  "127.0.0.1",
				Port:      3306,
				Parameter: "?parseTime=true",
			},
		},
		Token: token.Info{
			File: "rsa/token.pem",
		},
	}

	c := `
    {
        "Database": {
            "Type": "MySQL",
            "MySQL": {
                "Username":  "root",
                "Password":  "",
                "Name":      "achievers",
                "Hostname":  "127.0.0.1",
                "Port":      3306,
                "Parameter": "?parseTime=true"
            }
        },
        "Token": {
            "File": "rsa/token.pem"
        }
    }`

	actualConfig, err := New([]byte(c))

	if err != nil {
		t.Fatalf("New config returned error: %v", err)
	}

	if !cmp.Equal(expectedConfig, actualConfig) {
		t.Fatalf("Config returned unexpected value:\nexpected %#v,\nactual %#v",
			expectedConfig, actualConfig)
	}
}
